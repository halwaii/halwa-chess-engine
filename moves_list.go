package main

import "math/bits"

// this struct will save all moves
type MoveList struct{
	// Move == uint16 
	// slice is dynamic array
	Moves []Move
	count int // count of total no. of moves
}

// function to append moves into list
func AddMove(ml *MoveList, move Move){
	// add new move into slice
	ml.Moves = append(ml.Moves, move)
	ml.count++
}

// move extraction logic
// this will extract moves and add to our list
func ExtractMoves(fromSqare int, moveBitboard uint64,flag uint16, list *MoveList){
	
	for moveBitboard != 0{
		// 1) find destination square using LSB's index
		toSquare := bits.TrailingZeros64(moveBitboard)

		// 2) encode the move
		encodeMove := EncodeMove(uint16(fromSqare),uint16(toSquare),flag)

		// 3) add the encoded move to our list
		AddMove(list, encodeMove)

		// 4) clear the LSB's 1 so that we can find other possible moves too
		// moveBitboard     = 	0010010000
		// moveBitboard - 1 = 	0010001111
		moveBitboard &= (moveBitboard - 1)
	}
}

// extract pawn moves
func ExtractWhitePawnmoves(shift int, moveBitboard uint64, flag uint16, list *MoveList){
	// calculate tosquare and fromSquare
	for moveBitboard != 0 {
		toSquare := bits.TrailingZeros64(moveBitboard)

		// reverse math to find formSquare
		fromSquare := toSquare - shift

		// encode the move
		encodeMove := EncodeMove(uint16(fromSquare), uint16(toSquare), flag)
		// add move to ourlist
		AddMove(list, encodeMove)

		// update the bitboard
		moveBitboard &= (moveBitboard - 1)
	}
}
// same for black just 1 change
func ExtractBlackPawnmoves(shift int, moveBitboard uint64, flag uint16, list *MoveList){
	// calculate tosquare and fromSquare
	for moveBitboard != 0 {
		toSquare := bits.TrailingZeros64(moveBitboard)

		// reverse math to find formSquare
		// here is the change
		fromSquare := toSquare + shift

		// encode the move
		encodeMove := EncodeMove(uint16(fromSquare), uint16(toSquare), flag)
		// add move to ourlist
		AddMove(list, encodeMove)

		// update the bitboard
		moveBitboard &= (moveBitboard - 1)
	}
}