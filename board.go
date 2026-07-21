package main

import(
	"fmt"
)

// piece indentification to restored capture piece in undo move
const (
	Emptypiece = 0
	Whitepawn = 1
	WhiteKnight = 2
	WhiteBishop = 3
	WhiteRook = 4
	WhiteQueen = 5
	WhiteKing = 6

	BlackPawn = 7
	BlackKnight = 8
	BlackBishop = 9
	BlackRook = 10
	BlackQueen = 11
	BlackKing = 12
)
// instead of making copy of past bitboard
// we make a struct which carries only important info
// undostate struct to save past state memory
type UndoState struct {
	MoveMade Move // which move is played
	Capturedpiece int // which piece was captured
	Enpassantsq int // which was last en passant square
	CastlingRights uint8 // previous castling rights (KQkq) 0000 1111 
	HalfMoveClock int // 50 moves draw rule
}
// gloabal constants - A_lane and H_lane

// Column:    H1  G1  F1  E1  D1  C1  B1  A1
// Bits:      1   1   1   1   1   1   1   0
// Row:    8   7   6   5   4   3   2   1
// Hexa:   FE  FE  FE  FE  FE  FE  FE  FE

// FE = 1111 1110
const notA_lane uint64 = 0xFEFEFEFEFEFEFEFE
const notH_lane uint64 = 0x7F7F7F7F7F7F7F7F

// already making mask to check if there are any piece b/w king and rook
const (
	whiteKingsidemask = (uint64(1)<<5) | (uint64(1)<<6)
	whiteQueensidemask = (uint64(1)<<1) | (uint64(1)<<2) | (uint64(1)<<3)
	blackKingsidemask = (uint64(1)<<61) | (uint64(1)<<62)
	blackQueensidemask = (uint64(1)<<57) | (uint64(1)<<58) | (uint64(1)<<59)
)

// make a struct for every piece type and color
// to store board state
type board struct{
	// here every piece is a bitboard itself
	// 6 for white
	WhitePawns uint64; 
	WhiteKing uint64;
	WhiteQueen uint64;
	WhiteBishop uint64;
	WhiteKnight uint64;
	WhiteRook uint64;
	// 6 for black
	BlackPawns uint64;
	BlackKing uint64;
	BlackQueen uint64;
	BlackBishop uint64;
	BlackKnight uint64;
	BlackRook uint64;
	// square of last pawn which double pushed
	// game state variables
	EnPassantSquare int
	CastlingRights uint8
	HalfMoveClock int

	// undo stack
	history []UndoState

	// this will tell whose move is it
	WhiteToMove bool

	// hashkey 
	HashKey uint64
}

func Occupiedsquares(b board) uint64{
	// we get all occupied spaces as 1
	// we will do not of this and get empty spaces as 1
	return b.BlackBishop | b.BlackKing | b.BlackKnight | b.BlackPawns | b.BlackQueen | b.BlackRook |
			b.WhiteBishop | b.WhiteKing | b.WhiteKnight | b.WhitePawns | b.WhiteQueen | b.WhiteRook 
}

func whitepieces(b board) uint64{
	return b.WhiteBishop | b.WhiteKing | b.WhiteKnight | b.WhitePawns | b.WhiteQueen | b.WhiteRook
}
func blackpieces(b board) uint64{
	return b.BlackBishop | b.BlackKing | b.BlackKnight | b.BlackPawns | b.BlackQueen | b.BlackRook
}

func Printboard(b board) {
	fmt.Println("current board")
	for row:=7; row>=0; row--{
		fmt.Printf("%v ",row+1)
		for col:=0; col<8; col++{
			square := row*8 + col;
			if (b.BlackPawns & (1 << uint64(square))) != 0{
				fmt.Printf("p ")
			} else if (b.WhitePawns & (1 << uint64(square))) != 0{
				fmt.Printf("P ")
			} else if (b.WhiteKing & (1 << uint64(square))) != 0{
				fmt.Printf("K ")
			} else if (b.WhiteBishop & (1 << uint64(square))) != 0{
				fmt.Printf("B ")
			} else if (b.WhiteKnight & (1 << uint64(square))) != 0{
				fmt.Printf("N ")
			} else if (b.WhiteRook & (1 << uint64(square))) != 0{
				fmt.Printf("R ")
			} else if (b.WhiteQueen & (1 << uint64(square))) != 0{
				fmt.Printf("Q ")
			} else if (b.BlackKing & (1 << uint64(square))) != 0{
				fmt.Printf("k ")
			} else if (b.BlackQueen & (1 << uint64(square))) != 0{
				fmt.Printf("q ")
			} else if (b.BlackBishop & (1 << uint64(square))) != 0{
				fmt.Printf("b ")
			} else if (b.BlackKnight & (1 << uint64(square))) != 0{
				fmt.Printf("n ")
			} else if (b.BlackRook & (1 << uint64(square))) != 0 {
				fmt.Printf("r ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}
	fmt.Printf("\n  A B C D E F G H\n\n")
}

func PrintBitboard(bitboard uint64) {
    for row := 7; row >= 0; row-- {
        fmt.Printf("%v ", row+1)
        for col := 0; col < 8; col++ {
            square := row*8 + col
            // if sqare bit is 1 then print 1
            if (bitboard & (1 << uint64(square))) != 0 {
                fmt.Printf("1 ")
            } else {
                fmt.Printf(". ")
            }
        }
        fmt.Println()
    }
    fmt.Println("  A B C D E F G H")
}