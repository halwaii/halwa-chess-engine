package main

// will remove piece from the current bitboard

// bit logic
// bitboard &= ^mask (AND NOT)
// clear the bit for specific piece from specific square

func RemovePiece(b *board, piece int, square int){
	mask := uint64(1) << uint64(square)

	switch piece{
	case Whitepawn:
		b.WhitePawns &= ^mask
	case WhiteKnight:
		b.WhiteKnight &= ^mask
	case WhiteBishop:
		b.WhiteBishop &= ^mask
	case WhiteRook:
		b.WhiteRook &= ^mask
	case WhiteQueen:
		b.WhiteQueen &= ^mask
	case WhiteKing:
		b.WhiteKing &= ^mask
	case BlackPawn:
		b.BlackPawns &= ^mask
	case BlackKnight:
		b.BlackKnight &= ^mask
	case BlackBishop:
		b.BlackBishop &= ^mask
	case BlackRook:
		b.BlackRook &= ^mask
	case BlackQueen:
		b.BlackQueen &= ^mask
	case BlackKing:
		b.BlackKing &= ^mask
	}
}