package main

// function to get score value of a move
// MVV-LVA -> most valuable viction and least valuable attacker
// pawn attacking queen
// alpha beta logic is based on finding best moves in start
// so we can prune rest of the useless branches of tree to make engine efficient

// captures should have extremely high score
// same for promotion
func ScoreMoves(b *board, move Move)int{
	score := 0
	flag := GetFlag(move)

	// 1) capture (mvv-lva)
	capturedPiece := GetPieceAt(b, int(GetTo(move)))

	if capturedPiece!=Emptypiece{
		movingPiece := GetPieceAt(b, int(GetFrom(move)))

		// find value of victim and attacker
		victimVal := GetPieceValue(capturedPiece)
		attackerVal := GetPieceValue(movingPiece)

		// every capture is better than normal move, so add 1million to score
		// mvv-lva formula
		score += 1000000 + (10*victimVal) - attackerVal
	} else if flag == EpCapture{
		// enpassant case
		score += 1000000 + (10*PawnVal) - PawnVal
	}

	// 2) promotions
	if flag >= KnightPromo && flag<=QueenPromoCap{
		// promotion is also very valuable
		score += 900000
	}
	return score
}