package main

import "math/bits"

// pre calculated distance from center to push enemy king
// to edges
var CenterDist = [64]int{
	6, 5, 4, 3, 3, 4, 5, 6,
	5, 4, 3, 2, 2, 3, 4, 5,
	4, 3, 2, 1, 1, 2, 3, 4,
	3, 2, 1, 0, 0, 1, 2, 3,
	3, 2, 1, 0, 0, 1, 2, 3,
	4, 3, 2, 1, 1, 2, 3, 4,
	5, 4, 3, 2, 2, 3, 4, 5,
	6, 5, 4, 3, 3, 4, 5, 6,
}

func abs(n int) int{
	if n<0{
		return -n
	}
	return n
}
// mop-up logic -> push enemy king to edge/corner
func evalMopUp(b *board, iswhite bool) int{
	score := 0

	winnerKingSquare := bits.TrailingZeros64(b.WhiteKing)
	loserKingSquare := bits.TrailingZeros64(b.BlackKing)

	if !iswhite{
		winnerKingSquare = bits.TrailingZeros64(b.BlackKing)
		loserKingSquare = bits.TrailingZeros64(b.WhiteKing)
	}

	// 1) push enemy king to edge -> bonus
	score += CenterDist[loserKingSquare] * 10

	// 2) take winning king towards losing king
	// manhattan distance
	winnerRow, winnerCol := winnerKingSquare/8, winnerKingSquare%8
	loserRow, loserCol := loserKingSquare/8, loserKingSquare%8

	kingsDist := abs(winnerRow-loserRow) + abs(winnerCol-loserCol)

	// maximum distance b/w kings is 14 -> a1 and h8
	// the less the distance more the score
	// greedy scoring by multiplying
	score += (14 - kingsDist) * 4
	return score
}