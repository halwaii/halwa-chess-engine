package main

import(
	"strings"
	"strconv"
)

// fen -> shows whole position of board in single string
// there are 6 parts
// standard notation

// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1

// it starts from top left of board
// / 		-> represents new rank
// numbers 	-> show empty space
// ' ' 		-> empty space shows new part is started
// w / b 	-> shows whose turn is it
// KQkq 	-> shows castling rights
// '-' 		-> shows enpassnt square
// 0/1 		-> shows 50 move rule and full move

// we will split the string into 6 parts using
// string.Split() -> makes an array of strings
// and then solve individually for each part
func ParserFEN(b *board, fen string){

	// we will seperate where there is space
	parts := strings.Split(fen," ")

	// string stars from top and left of board
	// top is row 8 and left is col 1
	row := 7 // index
	col := 0

	for i:=0; i<len(parts[0]);i++{
		// take out each char
		char := parts[0][i]

		if char == '/' {
			// start new row
			row--
			// and restore col again to 0
			col = 0
		} else if char >= '1' && char <= '8'{
			// if numbers then then skip that many empty squres
			// char - '0' to convert into int
			emptySquares := int(char - '0')
			col += emptySquares
		} else {
			// if piece is present find its exact square
			// sqaure = row * 8 + col
			square :=  (row * 8) + col

			// convert char to piece 
			piece := CharToPiece(char)
			// add the piece to current square
			AddPiece(b, piece, square)
			// move to next square
			col++
		}
	}
	// turn
	if len(parts) > 1 {
		// if parts[1] is "w" then its white's turn
		b.WhiteToMove = (parts[1] == "w")
	}
	// castling rights
	if len(parts) > 2 {
		// at 1st no one can castle
		b.CastlingRights = 0
		if parts[2] != "-" {
			for i:=0;i<len(parts[2]);i++ {
				switch parts[2][i] {
				case 'K' : b.CastlingRights |= 1
				case 'Q' : b.CastlingRights |= 2
				case 'k' : b.CastlingRights |= 4
				case 'q' : b.CastlingRights |= 8
				}
			}
		}
	}
	// en passant square -> e4 ka hai e3
	b.EnPassantSquare = -1 // default
	if len(parts) > 3 && parts[3] != "-" {
		// e4 -> e is col and 4 is row
		// parts[3][0] is col ('a'-'h')
		// parts[3][1] is row ('1'-'8')
		epCol := int(parts[3][0] - 'a')
		epRow := int(parts[3][1] - '1')
		// formula to find exact square
		b.EnPassantSquare = (epRow*8) + epCol
	}
	// half move rule
	if len(parts) > 4 {
		halfMove, _ := strconv.Atoi(parts[4])
		b.HalfMoveClock = halfMove
	}
}