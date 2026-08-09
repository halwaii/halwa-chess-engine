package main

import "math/bits"

const (
	knightMobility = 4
	bishopMobility = 5
	rookMobility   = 2
	queenMobility  = 1
)

func evalMobility(b *board) int {
	return mobility(b, true) - mobility(b, false)
}

func mobility(b *board, isWhite bool) int {

	knight, bishop, rook, queen := b.WhiteKnight, b.WhiteBishop, b.WhiteRook, b.WhiteQueen
	if !isWhite {
		knight, bishop, rook, queen = b.BlackKnight, b.BlackBishop, b.BlackRook, b.BlackQueen
	}
	// create a copy of board to calculate score =  legal moves * moblility
	bitboard := *b
	score := 0

	for n := knight; n != 0; n &= n-1 {
		square := bits.TrailingZeros64(n)
		score += bits.OnesCount64(LegalKnightmoves(bitboard, square, isWhite)) * knightMobility
	}
	for n := bishop; n != 0; n &= n-1 {
		square := bits.TrailingZeros64(n)
		score += bits.OnesCount64(LegalBishopmoves(bitboard, square, isWhite)) * bishopMobility
	}
	for n := rook; n != 0; n &= n-1 {
		square := bits.TrailingZeros64(n)
		score += bits.OnesCount64(LegalRookmoves(bitboard, square, isWhite)) * rookMobility
	}
	for n := queen; n != 0; n &= n-1 {
		square := bits.TrailingZeros64(n)
		moves := LegalBishopmoves(bitboard, square, isWhite) | LegalRookmoves(bitboard, square, isWhite)
		score += bits.OnesCount64(moves) * queenMobility
	}
	return  score
}