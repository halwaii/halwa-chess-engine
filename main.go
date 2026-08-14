package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// helper function which converts index(0-63) to a1,e4,etc
func indexToSq(square uint16)string{
	row := square/8 // row 1 to 8
	col := square%8 // col a to h

	rowChar := rune('1'+row)
	colChar := rune('a'+col)

	return string(colChar)+string(rowChar)
}
// make move (uint16) to UCI string
func MoveToUCIstring(m Move) string{
	fromSq := GetFrom(m) // finding from square
	tosq := GetTo(m)
	flag := GetFlag(m)

	// ex - e2 to e4
	uci := indexToSq(fromSq)+indexToSq(tosq)

	// promotion
	switch flag{
	case QueenPromo, QueenPromoCap:
		uci += "q"
	case RookPromo, RookPromoCap:
		uci += "r"
	case BishopPromo, BishopPromoCap:
		uci += "b"
	case KnightPromo, KnightPromoCap:
		uci += "n"
	}
	return uci
}

// UCI string to move(uint16) conversion function
func UCIstringToMove(b *board, movestr string)Move{
	// generate all moves and then compare string with prev func
	var list MoveList
	GenerateAllMoves(b, &list)
	for i:=0;i<len(list.Moves);i++{
		move := list.Moves[i]
		if MoveToUCIstring(move)==movestr{
			return move
		}
	}
	return 0
}
// uci implementation
func UCILoop(){
	// gui gives command . reader -> reads that command
	reader := bufio.NewScanner(os.Stdin)
	var b board

	// zobrist should be initailized 1st
	InitZobrist()
	// initialize map with big memory
	TranspositionTable = make(map[uint64]TTEntry, 1000000)

	// main infinite loop till program ends
	for reader.Scan(){
		// take commands 
		command := strings.TrimSpace(reader.Text())
		// tokenization of string
		// "go depth 5" = ["go","depth","5"]
		tokens := strings.Fields(command)

		// to ignore empty enter
		if len(tokens) == 0{
			continue
		}

		switch tokens[0]{
		// checks if the engine understands uci commands
		case "uci":
			fmt.Println("id name Halwa_Engine")
			fmt.Println("id author Jay")
			fmt.Println("uciok")
		// keeps checking if engine is ok or stuck
		case "isready":
			fmt.Println("readyok")
		// new map for new game
		case "ucinewgame":
			TranspositionTable = make(map[uint64]TTEntry,1000000)

		case "position":
			// position startpos moves e2e4 e7e5
			// finds on which index is word "move"
			moveIdx := -1
			// 1) board setup
			if len(tokens)>1 && tokens[1]=="startpos"{
				ParserFEN(&b, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

				// check for "moves" word in array
				for i:=2;i<len(tokens);i++{
					if tokens[i]=="moves"{
						moveIdx=i+1// just next words are actual moves
						break
					}
				}
			} else if len(tokens)>=3 && tokens[1]=="fen"{
				fenEnd := len(tokens)
				for i:=2;i<len(tokens);i++{
					if tokens[i]=="moves"{
						moveIdx=i+1
						fenEnd=i
						break
					}
				}
				// again join fen string
				fenString := strings.Join(tokens[2:fenEnd]," ")
				ParserFEN(&b, fenString)
			}
			// 2) now Make move
			if moveIdx !=-1 && moveIdx<len(tokens){
				for i:=moveIdx; i<len(tokens); i++ {
					parsedMove := UCIstringToMove(&b, tokens[i])
					if parsedMove!=0{
						MakeMove(&b, parsedMove)
					} else {
						fmt.Println("info string ERROR: Move not found or illegal ->", tokens[i])
					}
				}
			}
			// 3) hash update after making move
			b.HashKey = GenerateHash(&b)

		case "go":
			// iterative deepeing
			var bestScore int
			for d:=1;d<=7;d++{
				bestScore = Search(&b, d, 0, -50000, 50000)
				fmt.Printf("info depth %d score cp %d\n", d, bestScore)
			}
			// how to find best move ?
			// search() returns bestscore
			// we save current best move in TT
			// so we use hashkey
			if entry, found := TranspositionTable[b.HashKey]; found {
				bestMove := entry.BestMove
				fmt.Printf("bestmove %s\n", MoveToUCIstring(bestMove))
			} else {
				fmt.Println("bestmove 0000")
			}
		case "quit":
			return
		}
	}
}
func main(){
	// UCILoop()
	StartWebServer()

	// var myMove Move = EncodeMove(12, 28, 0)
	// fmt.Println("pawn moves from e2(12) to e4(28) : \n")
	// fmt.Println("Encoded Move Value (Integer):", myMove)
	// fmt.Println("From Square:", myMove.GetFrom())
	// fmt.Println("To Square:", myMove.GetTo())
	// fmt.Println("Move Flag:", myMove.GetFlag())

	// var b board
	// b.WhitePawns = 0x000000000000FF00
	// b.WhiteKing = 0x0000000000000010
	// b.WhiteQueen = 0x0000000000000008
	// b.WhiteBishop = 0x0000000000000024 
	// b.WhiteKnight = 0x0000000000000042
	// b.WhiteRook = 0x0000000000000081
	
	// b.BlackPawns = 0x00FF000000000000
	// b.BlackKing = 0x1000000000000000
	// b.BlackRook = 0x8100000000000000
	// b.BlackBishop = 0x2400000000000000
	// b.BlackKnight = 0x4200000000000000
	// b.BlackQueen = 0x0800000000000000

	//Printboard(b)

	// fen testing ->
	//fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	// InitZobrist()					// 1) dictionary of random numbers
	// initialise new map
	// TranspositionTable = make(map[uint64]TTEntry, 1000000)
	//ParserFEN(&b, fen) 				// 2) set board pieces using fen string
	//b.HashKey = GenerateHash(&b)    // 3) now make hash of board
	//fmt.Printf("initial zobrist hash : 0x%v\n", b.HashKey)
	// Printboard(b)
	// perft

	// iterative deepening loop
	//startTime := time.Now()
	//var bestScore int

	// for d:=1;d<=6;d++{
	// 	bestScore = Search(&b, d, -50000,50000)
	// 	timeTaken := time.Since(startTime)
	// 	fmt.Printf("depth %d , Score: %d , time so far: %v\n",d,bestScore,timeTaken)
	// }
	//nodes := Perft(&b, 5)
	
	//fmt.Printf("\nTotal nodes for Depth 5 : %v \n", nodes)
	//fmt.Printf("\nFinal Best Evaluation Score : %v cp \n", bestScore)
	//fmt.Printf("Total time taken for depth 6 : %v \n", time.Since(startTime))

	// perft nodes calculation

	// startTime := time.Now()
	// nodes := Perft(&b, 6)
	// //perftDivide(&b, 6)
	// timeTaken := time.Since(startTime)
	// fmt.Printf("total nodes for depth 6 : %v \n", nodes)
	// fmt.Printf("total time taken : %v \n", timeTaken)

	// b.WhiteQueen = uint64(1) << 28
	// b.BlackPawns = uint64(1) << 46
	// fmt.Println("\nwhite queen legal moves(d4) and a blackpiece at g6 : \n")
	// queenmoves := allLegalQueenmoves(b, true)
	// PrintBitboard(queenmoves)
	// rookmoves := allLegalBlackRookmoves(b)
	// fmt.Println("\nblack Rook Legal Moves when it is at d4 and h8 and there is black piece(f4) and a white piece(d6):\n")
	// PrintBitboard(rookmoves)
	// whiteknightmoves := allLegalWhiteKnightmoves(b)
	// blackknightmoves := allLegalBlackKnightmoves(b)
	// fmt.Println("White Knight Legal Moves Matrix:")
	// PrintBitboard(whiteknightmoves)

	// fmt.Println("Black Knight Legal Moves Matrix:")
	// PrintBitboard(blackknightmoves)
}