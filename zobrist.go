package main

import (
	"math/rand"
	"time"
)

// to implement memoization we need to save the current board
// position as a key to use in hashmap

// we can't store it in FEN string (as it is slow)

// so we use zobrist hashing, which stores whole state of board
// in single 64-bit integer .

// we fill arrays with random numbers to prevent hash collision
// that is to prevent same number of 2 diff boards

// all these random numbers needs to be fixed at the start of engine
// as they act as key in our transposition table

// 1) global zobrist table
// i wrote 13 cause 0 was emptypiece
var ZobristPieces [13][64]uint64 // 12 pieces and 64 squares
var ZobristCastling [16]uint64 // castling rights
var ZobristEnpassant [8]uint64 // 8 cols for enpassnt (a to h)
var ZobristTurn uint64 // 1 random number to toggle turn

// generate random 64 bit number
func randomUint64() uint64{
	// to create true randomness we use this logic
	return uint64(rand.Uint32())<<32 | uint64(rand.Uint32())
}

// 2) zobrist function
func InitZobrist(){
	// to create diff random numbers we seed it with curr time
	rand.Seed(time.Now().UnixNano())

	// 1) fill random numbers for pieces
	for piece :=0 ; piece<13; piece++{
		for square := 0; square<64;square++{
			ZobristPieces[piece][square] = randomUint64()
		}
	}

	// 2) random no. for castling rights
	for i:=0;i<16;i++{
		ZobristCastling[i] = randomUint64()
	}

	// 3) enpassnt cols
	for i:=0;i<8;i++{
		ZobristEnpassant[i] = randomUint64()
	}

	// 4) toggle turn
	ZobristTurn = randomUint64()
}

// 3) generate an unique zobrist hash from board state
// to do this we will do XOR of random numbers
func GenerateHash(b *board) uint64{
	var finalHash uint64 = 0

	// 1) check all 64 squares and XOR the random number where piece is present
	for square:=0; square<64;square++{
		piece := GetPieceAt(b, square)
		if piece != Emptypiece{
			finalHash ^= ZobristPieces[piece][square]
		}
	}

	// 2) check turn
	if b.WhiteToMove {
		finalHash ^= ZobristTurn
	}
	
	// 3) castling rights
	finalHash ^= ZobristCastling[b.CastlingRights]

	// 4) en passant 
	if b.EnPassantSquare != -1{
		// to find column from 0 to 7
		col := b.EnPassantSquare % 8
		finalHash ^= ZobristEnpassant[col]
	}

	return finalHash
}