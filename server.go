package main

import (
	"fmt"
	"log"
	"net/http"
	"sync" // NEW

	"github.com/gorilla/websocket"
)

// In-Memory Stats Tracker
type EngineStats struct {
	Visits     int `json:"visits"`
	Games      int `json:"games"`
	EngineWins int `json:"engineWins"`
	Draws      int `json:"draws"`
}
var globalStats EngineStats

// NEW: protects Search()/TranspositionTable - only one game thinks at a time
var engineMutex sync.Mutex

// NEW: protects globalStats, since multiple WebSocket connections
// (i.e. multiple visitors) can hit this at the same time
var statsMutex sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WSMessage struct {
	Fen    string `json:"fen"`
	Result string `json:"result"` // "draw" or "enginewin"
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// NEW: lock around the stats read/write + broadcast
	statsMutex.Lock()
	globalStats.Visits++
	conn.WriteJSON(globalStats)
	statsMutex.Unlock()

	for {
		var msg WSMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}

		if msg.Result != "" {
			// NEW: lock around this whole stats-mutation block
			statsMutex.Lock()
			globalStats.Games++
			if msg.Result == "enginewin" {
				globalStats.EngineWins++
			} else if msg.Result == "draw" {
				globalStats.Draws++
			}
			conn.WriteJSON(globalStats)
			statsMutex.Unlock()
			continue
		}

		if msg.Fen != "" {
			fmt.Println("Received FEN:", msg.Fen)

			var b board
			ParserFEN(&b, msg.Fen)
			b.HashKey = GenerateHash(&b)

			// NEW: lock around Search + TT lookup - critical fix,
			// without this two simultaneous games can crash the server
			engineMutex.Lock()
			Search(&b, 5, 0, -50000, 50000)
			bestMoveStr := "0000"
			if entry, found := TranspositionTable[b.HashKey]; found {
				bestMoveStr = MoveToUCIstring(entry.BestMove)
			}
			engineMutex.Unlock()

			conn.WriteJSON(map[string]string{"move": bestMoveStr})
		}
	}
}

func StartWebServer() {
	InitZobrist()
	TranspositionTable = make(map[uint64]TTEntry, 1000000)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/ws", handleWebSocket)

	fmt.Println("🚀 Halwa Engine Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}