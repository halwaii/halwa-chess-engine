package main

import "math/bits"

const (
	// penalty for openfile, semi openfile, no pawn shield
	OpenFile     = -25
	SemiOpenFile = -13
	NoPawnShield = -10
)

func evalKingSafety(b *board) int {
	// white safety - black safety
	return kingSafety(b, true) - kingSafety(b, false)
}

func kingSafety(b *board, iswhite bool) int {
	kingBit, ownPawns, enemyPawns := b.WhiteKing, b.WhitePawns, b.BlackPawns

	if !iswhite{
		kingBit, ownPawns, enemyPawns = b.BlackKing, b.BlackPawns, b.WhitePawns
	}
	square := bits.TrailingZeros64(kingBit)
	col := square%8
	row := square/8
	score := 0

	// open (no pawns from both sides)
	// and semi open (pawn of enemy side only) file check
	own := ownPawns & FileMasks[col]
	enemy := enemyPawns & FileMasks[col]

	if own == 0 && enemy == 0{
		score += OpenFile // no pawns from both sides
	} else if own == 0{
		score += SemiOpenFile
	}

	// pawn shield check
	// king should be shielded by 3 pawns in 3 files
	shieldRow := row + 1
	if !iswhite{
		shieldRow = row - 1
	}
	if shieldRow >=0 && shieldRow <8{
		// check 3 files -> left, center, right (-1, 0 , 1)
		for c:= col-1; c <= col+1;c++{
			if c<0 || c>7{
				continue // ignore out of bound files
			}
			shieldSquare := shieldRow * 8 + c
			// if missing then penalty
			if (own & (uint64(1)<<uint64(shieldSquare))) == 0 {
				score += NoPawnShield
			}
		}
	}
	return score
}