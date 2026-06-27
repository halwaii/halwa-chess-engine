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

func Blackpawnpush(b board, occupied uint64, whitepieces uint64) uint64{
	// single push
	Singlepush := (b.BlackPawns >> 8) & ^occupied

	// double push
	// row 6 and 5 should be empty
	mask6row := uint64(0x0000FF0000000000)
	Doublepush := ((Singlepush & mask6row) >> 8) & ^occupied

	// capture
	// capture can be left (>> 9) and not on A-lane
	Leftcapture := ((b.BlackPawns & notA_lane) >> 9) & whitepieces
	// capture can be right (>> 7) and not on H-lane
	Rightcapture := ((b.BlackPawns & notH_lane) >> 7) & whitepieces

	return Singlepush | Doublepush | Leftcapture | Rightcapture
}
func whitepieces(b board) uint64{
	return b.WhiteBishop | b.WhiteKing | b.WhiteKnight | b.WhitePawns | b.WhiteQueen | b.WhiteRook
}
func blackpieces(b board) uint64{
	return b.BlackBishop | b.BlackKing | b.BlackKnight | b.BlackPawns | b.BlackQueen | b.BlackRook
}

// all possible knight moves
func Knightmoves(square int) uint64{
	// create moves to store all possible moves of knight
	var moves uint64 = 0
	// first find row and column of our knight
	row := square / 8
	col := square % 8

	// all 8 - L shaped possible offsets of knight
	rowOffsets := []int{2,2,1,1,-1,-1,-2,-2}
	colOffsets := []int{1,-1,2,-2,2,-2,1,-1}

	for i:=0 ; i<8; i++ {
		// finding rows and columns of possible moves
		r := row + rowOffsets[i]
		c := col + colOffsets[i]
		
		// edge cases
		// possible moves should not go beyond
		if r>=0 && r<8 && c>=0 && c<8 {
			PossibleSquare := r*8 + c
			moves |= ((uint64(1) << uint64(PossibleSquare)))
		}
	}
	return moves
}
// all Legal white Knight moves
func LegalKnightmoves(b board, square int, isWhite bool) uint64{
	rawmoves := Knightmoves(square)
	// so that white knight cannnot take white pieces
	if isWhite {
		return rawmoves & ^whitepieces(b)
	} else {
		return rawmoves & ^blackpieces(b)
	}
}
// all Legal Black Knight moves
// func LegalBlackKnightmoves(b board, square int) uint64{
// 	rawmoves := Knightmoves(square)
// 	// so that black knight cannnot take black pieces
// 	return rawmoves & ^blackpieces(b)
// }
// we will find the square on which our knight is actually standing
func allLegalKnightmoves(b board, isWhite bool) uint64{
	var allKnightmoves uint64 = 0
	for square := 0 ; square < 64 ; square++ {
		var currKnight uint64 = 0
		if isWhite {
			currKnight = b.WhiteKnight
		} else {
			currKnight = b.BlackKnight
		}
		if (currKnight & (uint64(1) << uint64(square))) != 0 {
			allKnightmoves |= LegalKnightmoves(b, square, isWhite)
		}
	}
	return allKnightmoves
}
// func allLegalBlackKnightmoves(b board) uint64{
// 	var allBlackKnightmoves uint64 = 0
// 	for square := 0 ; square < 64 ; square++ {
// 		if (b.BlackKnight & (uint64(1) << uint64(square))) != 0 {
// 			allBlackKnightmoves |= LegalBlackKnightmoves(b, square)
// 		}
// 	}
// 	return allBlackKnightmoves
// }

// all possible sliding moves
func Slidingmoves(square int, b board, rowOffsets, colOffsets []int) uint64{
	var moves uint64 = 0
	// find row and col of current rook/bishop/queen
	row := square / 8
	col := square % 8

	// ray casting in every direction
	// here instead of 4 we will do len(rowOffsets) 
	// cause queen can move in 8 directions
	for i :=0;i< len(rowOffsets) ;i++ {
		r := row + rowOffsets[i]
		c := col + colOffsets[i]
		// going on till end of board
		for r>=0 && r<8 && c>=0 && c<8 {
			PossibleSquare := r*8 + c
			mask := (uint64(1) << uint64(PossibleSquare))
			moves |= mask
			// break considtion if a square is occupied
			if (Occupiedsquares(b) & mask) != 0 {
				break
			}
			// increment/decrement
			r += rowOffsets[i]
			c += colOffsets[i]
		}
	}
	return moves
}
// instead of writing different functions for black and white we wrote in single funciton
// get all Legal rook moves
func LegalRookmoves(b board, square int, isWhite bool) uint64{
	// rook can move in EAST WEST NORTH SOUTH directions
	rowOffsets := []int{1,-1,0,0}
	colOffsets := []int{0,0,1,-1}
	rawmoves := Slidingmoves(square, b, rowOffsets, colOffsets)

	if isWhite {
		return rawmoves & ^whitepieces(b)
	} else {
		return rawmoves & ^blackpieces(b)
	}
}
// find squares where rooks are present
func allLegalRookmoves(b board, isWhite bool) uint64{
	var allRookmoves uint64 = 0
	for square:=0 ; square<64; square++{
		var currentRook uint64 = 0
		if isWhite {
			currentRook = b.WhiteRook
		} else {
			currentRook = b.BlackRook
		}
		if (currentRook & (uint64(1)<<uint64(square))) != 0{
			allRookmoves |= LegalRookmoves(b, square, isWhite)
		}
	}
	return allRookmoves
}
// same logic is for bishops
func LegalBishopmoves(b board, square int, isWhite bool) uint64{
	// bishop moves in NE , NW , SE , SW
	rowOffsets := []int{1,1,-1,-1}
	colOffsets := []int{1,-1,1,-1}
	
	rawmoves := Slidingmoves(square, b, rowOffsets, colOffsets)

	if isWhite {
		return rawmoves & ^whitepieces(b)
	} else {
		return rawmoves & ^blackpieces(b)
	}
}
func allLegalBishopmoves(b board, isWhite bool) uint64{
	var allBishopmoves uint64 = 0
	for square := 0 ; square < 64 ; square++ {
		var currentBishop uint64 = 0
		if isWhite {
			currentBishop = b.WhiteBishop
		} else {
			currentBishop = b.BlackBishop
		}
		if (currentBishop & (uint64(1)<<uint64(square))) != 0 {
			allBishopmoves |= LegalBishopmoves(b, square, isWhite)
		}
	}
	return allBishopmoves
}
// queen = bishop + rook
// it can move in 8 directions
func allLegalQueenmoves(b board, isWhite bool) uint64{
	var allQueenmoves uint64 = 0
	for square :=0 ; square<64; square++ {
		var currentQueen uint64 = 0
		if isWhite {
			currentQueen = b.WhiteQueen
		} else {
			currentQueen = b.BlackQueen
		}
		if (currentQueen & (uint64(1)<<uint64(square))) != 0{
			allQueenmoves |= LegalBishopmoves(b, square, isWhite) | LegalRookmoves(b, square, isWhite)
		}
	}
	return allQueenmoves
}
// all possisble king moves
func Kingmoves(square int) uint64{
	var moves uint64 = 0
	// find row and col of king's position
	row := square / 8
	col := square % 8
	// offsets for 8 directions
	rowOffsets := []int{1,1,1,0,0,-1,-1,-1}
	colOffsets := []int{-1,0,1,1,-1,-1,0,1}

	for i:=0; i<8; i++{
		// next row and cols
		r := row + rowOffsets[i]
		c := col + colOffsets[i]

		if r>=0 && r<8 && c>=0 && c<8 {
			// target square
			PossibleSquare := r*8 + c
			moves |= (uint64(1)<<uint64(PossibleSquare))
		}
	}
	return moves
}
func LegalKingmoves(b board, square int, isWhite bool) uint64{
	rawmoves := Kingmoves(square)

	if isWhite {
		return rawmoves & ^whitepieces(b)
	} else {
		return rawmoves & ^blackpieces(b)
	}
}
func allLegalKingmoves(b board, isWhite bool) uint64{
	var allKingmoves uint64 = 0
	for square:=0; square<64; square++ {
		var currKing uint64 = 0;
		if isWhite{
			currKing = b.WhiteKing
		} else {
			currKing = b.BlackKing
		}
		if (currKing & (uint64(1)<<uint64(square))) != 0{
			allKingmoves |= LegalKingmoves(b, square, isWhite)
		}
	}
	return allKingmoves
}
// func LegalBlackRookmoves(b board, square int) uint64{
// 	rowOffsets := []int{1,-1,0,0}
// 	colOffsets := []int{0,0,1,-1}
// 	rawmoves := Slidingmoves(square, b, rowOffsets, colOffsets)
// 	// filter so that they can't take their friendly piece
// 	return rawmoves & ^blackpieces(b)
// }

// func allLegalBlackRookmoves(b board) uint64{
// 	var allRookmoves uint64 = 0
// 	for square:=0 ; square<64 ; square++ {
// 		if (b.BlackRook & (uint64(1) << uint64(square))) != 0{
// 			allRookmoves |= LegalBlackRookmoves(b, square)
// 		}
// 	}
// 	return allRookmoves
// }
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
    fmt.Println("  A B C D E F G H\n")
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