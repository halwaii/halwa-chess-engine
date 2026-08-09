package main

import "math/bits"

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
	KingVal = 20000
)

// PST -> every piece on every square has a specific bonus or penalty
// so we define a PST array for each piece

// pawns have high values on high rank
var PawnMiddlegamePST = [64]int{
	0,0,0,0,0,0,0,0, // pawns dont exist here
	60,60,60,60,60,60,60,60, // rank 7 most valuable due to promotion
	10,10,20,30,30,20,10,10,
	5,5,10,45,45,10,5,5,
	0,0,0,40,40,0,0,0,
	5,-5,-10,0,0,-10,-5,5,
	5,10,10,-20,-20,10,10,5,
	0,0,0,0,0,0,0,0,
}
// pst for end game
var PawnEndgamePST = [64]int{
	0,  0,  0,  0,  0,  0,  0,  0,
	 80, 80, 80, 80, 80, 80, 80, 80, // promotion
	 50, 50, 50, 50, 50, 50, 50, 50, 
	 30, 30, 30, 30, 30, 30, 30, 30, 
	 20, 20, 20, 20, 20, 20, 20, 20,
	 10, 10, 10, 10, 10, 10, 10, 10,
	 10, 10, 10, 10, 10, 10, 10, 10,
	  0,  0,  0,  0,  0,  0,  0,  0,
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
var KingMiddlegamePST = [64]int{
	-30,-40,-40,-50,-50,-40,-40,-30,
	-30,-40,-40,-50,-50,-40,-40,-30,
	-30,-40,-40,-50,-50,-40,-40,-30,
	-30,-40,-40,-50,-50,-40,-40,-30,
	-20,-30,-30,-40,-40,-30,-30,-20,
	-10,-20,-20,-20,-20,-20,-20,-10,
	 20, 20,  0,  0,  0,  0, 20, 20,
	 20, 30, 10,  0,  0, 10, 30, 20,
}

// king is stronger in middle during endgames
var KingEndgamePST = [64]int{
	-50,-30,-30,-30,-30,-30,-30,-50,
	-30,-20,  0,  0,  0,  0,-20,-30,
	-30,  0, 20, 30, 30, 20,  0,-30,
	-30,  0, 30, 40, 40, 30,  0,-30,
	-30,  0, 30, 40, 40, 30,  0,-30,
	-30,  0, 20, 30, 30, 20,  0,-30,
	-30,-20,  0,  0,  0,  0,-20,-30,
	-50,-30,-30,-30,-30,-30,-30,-50,
}

// game phase constants
const (
	knightPhase = 1
	bishopPhase = 1
	rookPhase = 2
	queenPhase = 4
	// 4 knights, 4 bishops, 4 rooks, 2 queens
	totalphase = 24
)
// tells engine if the game is in middle game or endgame
func GamePhase(b *board) int{
	phase := totalphase
	phase -= bits.OnesCount64(b.WhiteKnight | b.BlackKnight) * knightPhase
	phase -= bits.OnesCount64(b.WhiteBishop | b.BlackBishop) * bishopPhase
	phase -= bits.OnesCount64(b.WhiteRook | b.BlackRook) * rookPhase
	phase -= bits.OnesCount64(b.WhiteQueen | b.BlackQueen) * queenPhase

	if phase<0{
		phase = 0
	}
	return  phase
}
// this funciton tells engine whose favour is in curr board
func Evaluate(b *board) int{
	mgScore, egScore := 0, 0

	// loop on all squares
	for i:=0;i<64;i++{
		piece := GetPieceAt(b, i)

		if piece != Emptypiece{
			// add for white
			switch piece{
			case Whitepawn: 
				mgScore += PawnVal + PawnMiddlegamePST[i]
				egScore += PawnVal + PawnEndgamePST[i]
			case WhiteKnight: 
				mgScore += KnightVal + KnightPSt[i]
				egScore += KnightVal + KnightPSt[i]
			case WhiteBishop: 
				mgScore += BishopVal + BishopPST[i]
				egScore += BishopVal + BishopPST[i]
			case WhiteRook: 
				mgScore += RookVal + RookPST[i]
				egScore += RookVal + RookPST[i]
			case WhiteQueen: 
				mgScore += QueenVal + QueenPST[i]
				egScore += QueenVal + QueenPST[i]
			case WhiteKing: 
				mgScore += KingVal + KingMiddlegamePST[i]
				egScore += KingVal + KingEndgamePST[i]

			// subtract for black
			case BlackPawn: 
				mgScore -= PawnVal + PawnMiddlegamePST[mirrorsquare(i)]
				egScore -= PawnVal + PawnEndgamePST[mirrorsquare(i)]
			case BlackKnight: 
				mgScore -= KnightVal + KnightPSt[mirrorsquare(i)]
				egScore -= KnightVal + KnightPSt[mirrorsquare(i)]
			case BlackBishop: 
				mgScore -= BishopVal + BishopPST[mirrorsquare(i)]
				egScore -= BishopVal + BishopPST[mirrorsquare(i)]
			case BlackRook: 
				mgScore -= RookVal + RookPST[mirrorsquare(i)]
				egScore -= RookVal + RookPST[mirrorsquare(i)]
			case BlackQueen: 
				mgScore -= QueenVal + QueenPST[mirrorsquare(i)]
				egScore -= QueenVal + QueenPST[mirrorsquare(i)]
			case BlackKing: 
				mgScore -= KingVal + KingMiddlegamePST[mirrorsquare(i)]
				egScore -= KingVal + KingEndgamePST[mirrorsquare(i)]
			}
		}
	}

	pawnMG, pawnEG := EvaluatePawnStructure(b)
	mgScore += pawnMG
	egScore += pawnEG
	mgScore += evalKingSafety(b)
	mgScore += evalDevelopment(b)
	mgScore += evalCastling(b)

	// add eval mop up according to end game
	// check who is winning
	if egScore > 100 {
		egScore += evalMopUp(b, true)
	} else if egScore < -100{
		egScore -= evalMopUp(b, false)
	}

	phase := GamePhase(b)
	blended := (mgScore*(24-phase) + egScore*phase) / 24

	mobilityScore := evalMobility(b)
	
	score := blended + mobilityScore 
	// return score according to curr player perspective
	if b.WhiteToMove{
		return score
	}
	return -score
}