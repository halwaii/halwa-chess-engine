package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections
	},
}

// Struct to receive data from frontend
type WSMessage struct {
	Fen string `json:"fen"`
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		// 1. Browser se FEN string read karo
		var msg WSMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		fmt.Println("Received FEN from browser:", msg.Fen)

		// 2. Halwa Engine ka board setup karo
		var b board
		ParserFEN(&b, msg.Fen) 
		b.HashKey = GenerateHash(&b) 

		// 3. Engine ko sochne bolo (Depth 5 for quick response)
		Search(&b, 5, 0, -50000, 50000) 

		// 4. TT (Transposition Table) se best move nikalo
		bestMoveStr := "0000"
		if entry, found := TranspositionTable[b.HashKey]; found { 
			bestMoveStr = MoveToUCIstring(entry.BestMove) 
		}
		fmt.Println("Halwa Engine plays:", bestMoveStr)

		// 5. Best move wapas frontend ko bhej do
		conn.WriteJSON(map[string]string{"move": bestMoveStr})
	}
}

// Function to start server
func StartWebServer() {
	// Engine initialization
	InitZobrist()
	TranspositionTable = make(map[uint64]TTEntry, 1000000)
	
	// Serve HTML/JS/CSS files from current directory
	http.Handle("/", http.FileServer(http.Dir(".")))
	
	// WebSocket endpoint setup
	http.HandleFunc("/ws", handleWebSocket)

	fmt.Println("🚀 Halwa Engine Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}