package main

// we measure this score in Centipawns
// cp = more +ve -> white is winning
// cp = more -ve -> black is winning

// evaluation is based on 2 things
// 1) material value (greed)
// 2) Piece square tables (positional sense)
// 
const(
	PawnVal = 100
	KnightVal = 300
	BishopVal = 320
	RookVal = 500
	QueenVal = 900
	KingVal = 50000
)

// this funciton tells engine whose favour is in curr board
func Evaluate(b *board) int{
	score := 0

	// loop on all squares
	for i:=0;i<64;i++{
		piece := GetPieceAt(b, i)

		if piece != Emptypiece{
			// add for white
			switch piece{
			case Whitepawn: 
				score += PawnVal
			case WhiteKnight: 
				score += KnightVal
			case WhiteBishop: 
				score += BishopVal
			case WhiteRook: 
				score += RookVal
			case WhiteQueen: 
				score += QueenVal
			case WhiteKing: 
				score += KingVal

			// subtract for black
			case BlackPawn: 
				score -= PawnVal
			case BlackKnight: 
				score -= KnightVal
			case BlackBishop: 
				score -= BishopVal
			case BlackRook: 
				score -= RookVal
			case BlackQueen: 
				score -= QueenVal
			case BlackKing: 
				score -= KingVal
			}
		}
	}
	return score
}