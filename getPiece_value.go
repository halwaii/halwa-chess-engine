package main

// helper function to get value of given piece
// which will be used for move ordering
func GetPieceValue(piece int)int{
	switch piece{
	case Whitepawn, BlackPawn:
		return PawnVal // 100
	case WhiteKnight, BlackKnight:
		return KnightVal // 300
	case WhiteBishop, BlackBishop:
		return BishopVal // 320
	case WhiteRook, BlackRook:
		return RookVal // 500
	case WhiteQueen, BlackQueen:
		return QueenVal // 900
	case WhiteKing, BlackKing:
		return KingVal // 20000
	default:
		return 0
	}
}