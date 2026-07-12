package main

// using this function we find what piece is present on that square
// and help to detect capture in make move
func GetPieceAt(b *board, square int) int {
	mask := uint64(1) << uint64(square)

	if(b.WhitePawns & mask) != 0 { return Whitepawn}
	if(b.WhiteKnight & mask) != 0 { return WhiteKnight}
	if(b.WhiteBishop & mask) != 0 { return WhiteBishop}
	if(b.WhiteRook & mask) != 0 { return WhiteRook}
	if(b.WhiteQueen & mask) != 0 { return WhiteQueen}
	if(b.WhiteKing & mask) != 0 { return WhiteKing}
	if(b.BlackPawns & mask) != 0 { return BlackPawn}
	if(b.BlackKnight & mask) != 0 { return BlackKnight}
	if(b.BlackBishop & mask) != 0 { return BlackBishop}
	if(b.BlackRook & mask) != 0 { return BlackRook}
	if(b.BlackQueen & mask) != 0 { return BlackQueen}
	if(b.BlackKing & mask) != 0 { return BlackKing}
	// if no piece is present return empty
	return Emptypiece
}