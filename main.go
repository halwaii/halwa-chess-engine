package main

import (
	"fmt"
)

// gloabal constants - A_lane and H_lane

// Column:    H1  G1  F1  E1  D1  C1  B1  A1
// Bits:      1   1   1   1   1   1   1   0
// Row:    8   7   6   5   4   3   2   1
// Hexa:   FE  FE  FE  FE  FE  FE  FE  FE

// FE = 1111 1110
const notA_lane uint64 = 0xFEFEFEFEFEFEFEFE
const notH_lane uint64 = 0x7F7F7F7F7F7F7F7F

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
}

func Printboard(b board) {
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

func Occupiedsquares(b board) uint64{
	// we get all occupied spaces as 1
	// we will do not of this and get empty spaces as 1
	return b.BlackBishop | b.BlackKing | b.BlackKnight | b.BlackPawns | b.BlackQueen | b.BlackRook |
			b.WhiteBishop | b.WhiteKing | b.WhiteKnight | b.WhitePawns | b.WhiteQueen | b.WhiteRook 
}

// move pawn
func Whitepawnpush(b board,occupied uint64, blackpieces uint64) uint64{
	// single push
	Singlepush := (b.WhitePawns << 8) & ^occupied 

	// double push
	// and row 3 and row 4 should be empty
	row3mask := uint64(0x000000000000FF0000)
	Doublepush := ((Singlepush & row3mask)<<8) & ^occupied

	// capture
	// capture can be left (<< 7) but it should not be on A-lane and it can only capture black piece
	Leftcapture := ((b.WhitePawns & notA_lane) << 7) & blackpieces

	// capture can be right (<< 9) but it should not be on H-lane and there should be black piece
	Rightcapture := ((b.WhitePawns & notH_lane) << 9) & blackpieces

	return Singlepush | Doublepush | Leftcapture |Rightcapture
}
func blackpieces(b board) uint64{
	return b.BlackBishop | b.BlackKing | b.BlackKnight | b.BlackPawns | b.BlackQueen | b.BlackRook
}
func main(){
	var b board
	b.WhitePawns = 0x000000000000FF00
	b.WhiteKing = 0x0000000000000010
	b.WhiteQueen = 0x0000000000000008
	b.WhiteBishop = 0x0000000000000024
	b.WhiteKnight = 0x0000000000000042
	b.WhiteRook = 0x0000000000000081
	
	b.BlackPawns = 0x00FF000000000000
	b.BlackKing = 0x1000000000000000
	b.BlackRook = 0x8100000000000000
	b.BlackBishop = 0x2400000000000000
	b.BlackKnight = 0x4200000000000000
	b.BlackQueen = 0x0800000000000000

	
	fmt.Println("current board\n")
	Printboard(b)
}