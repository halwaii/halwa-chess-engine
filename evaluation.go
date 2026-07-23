package main

// we measure this score in Centipawns
// cp = more +ve -> white is winning
// cp = more -ve -> black is winning

// evaluation is based on 2 things
// 1) material value (greed)
// 2) Piece square tables (positional sense)

// by doing "^ 56" we can vertically flip the board
// and assign values for black pieces too instead of writing whole code again
func mirrorsquare(sq int) int{
	return sq ^ 56
}
const(
	PawnVal = 100
	KnightVal = 300
	BishopVal = 320
	RookVal = 500
	QueenVal = 900
	KingVal = 50000
)

// PST -> every piece on every square has a specific bonus or penalty
// so we define a PST array for each piece

// pawns have high values on high rank
var PawnPST = [64]int{
	0,0,0,0,0,0,0,0, // pawns dont exist here
	50,50,50,50,50,50,50,50, // rank 7 most valuable due to promotion
	10,10,20,30,30,20,10,10,
	5,5,10,25,25,10,5,5,
	0,0,0,20,20,0,0,0,
	5,-5,-10,0,0,-10,-5,5,
	5,10,10,-20,-20,10,10,5,
	0,0,0,0,0,0,0,0,
}

// knights have high values in center and low on corner/edges
var KnightPSt = [64]int{
	-50,-40,-30,-30,-30,-30,-40,-50,
	-40,-20,0,0,0,0,-20,-40,
	-30,0,10,15,15,10,0,-30,
	-30, 5,15,20,20, 15,5,-30,
	-30,0, 15, 20, 20, 15,0,-30,
	-30,5, 10,15,15,10,5,-30,
	-40,-20,0,5,5,0,-20,-40,
	-50,-40,-30,-30,-30,-30,-40,-50,
}

// rooks prefers open files and 7th rank (pawns ka)
var RookPST = [64]int{
	0,  0,  0,  0,  0,  0,  0,  0,
	5, 10, 10, 10, 10, 10, 10,  5,
	-5,  0,  0,  0,  0,  0,  0, -5,
	-5,  0,  0,  0,  0,  0,  0, -5,
	-5,  0,  0,  0,  0,  0,  0, -5,
	-5,  0,  0,  0,  0,  0,  0, -5,
	-5,  0,  0,  0,  0,  0,  0, -5,
	0,  0,  0,  5,  5,  0,  0,  0,
}

// bishop wants to control long diagonals and center
var BishopPST = [64]int{
	-20,-10,-10,-10,-10,-10,-10,-20,
	-10,  0,  0,  0,  0,  0,  0,-10,
	-10,  0,  5, 10, 10,  5,  0,-10,
	-10,  5,  5, 10, 10,  5,  5,-10,
	-10,  0, 10, 10, 10, 10,  0,-10,
	-10, 10, 10, 10, 10, 10, 10,-10,
	-10,  5,  0,  0,  0,  0,  5,-10,
	-20,-10,-10,-10,-10,-10,-10,-20,
}

// queen is mix
var QueenPST = [64]int{
	-20,-10,-10, -5, -5,-10,-10,-20,
	-10,  0,  0,  0,  0,  0,  0,-10,
	-10,  0,  5,  5,  5,  5,  0,-10,
	 -5,  0,  5,  5,  5,  5,  0, -5,
	  0,  0,  5,  5,  5,  5,  0, -5,
	-10,  5,  5,  5,  5,  5,  0,-10,
	-10,  0,  5,  0,  0,  0,  0,-10,
	-20,-10,-10, -5, -5,-10,-10,-20,
}

// king wants to hide in corners (midgame) behind pawns
var KingPST = [64]int{
	-30,-40,-40,-50,-50,-40,-40,-30,
	-30,-40,-40,-50,-50,-40,-40,-30,
	-30,-40,-40,-50,-50,-40,-40,-30,
	-30,-40,-40,-50,-50,-40,-40,-30,
	-20,-30,-30,-40,-40,-30,-30,-20,
	-10,-20,-20,-20,-20,-20,-20,-10,
	 20, 20,  0,  0,  0,  0, 20, 20,
	 20, 30, 10,  0,  0, 10, 30, 20,
}
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
				score += PawnPST[i]
			case WhiteKnight: 
				score += KnightVal
				score += KnightPSt[i]
			case WhiteBishop: 
				score += BishopVal
				score += BishopPST[i]
			case WhiteRook: 
				score += RookVal
				score += RookPST[i]
			case WhiteQueen: 
				score += QueenVal
				score += QueenPST[i]
			case WhiteKing: 
				score += KingVal
				score += KingPST[i]

			// subtract for black
			case BlackPawn: 
				score -= PawnVal
				score -= PawnPST[mirrorsquare(i)]
			case BlackKnight: 
				score -= KnightVal
				score -= KnightPSt[mirrorsquare(i)]
			case BlackBishop: 
				score -= BishopVal
				score -= BishopPST[mirrorsquare(i)]
			case BlackRook: 
				score -= RookVal
				score -= RookPST[mirrorsquare(i)]
			case BlackQueen: 
				score -= QueenVal
				score -= QueenPST[mirrorsquare(i)]
			case BlackKing: 
				score -= KingVal
				score -= KingPST[mirrorsquare(i)]
			}
		}
	}
	// return score according to curr player perspective
	if b.WhiteToMove{
		return score
	}
	return -score
}