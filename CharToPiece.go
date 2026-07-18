package main

// helper funciton to map FEN char to pieces
func CharToPiece(c byte) int{
	switch c {
	case 'P' : return Whitepawn
	case 'N' : return WhiteKnight
	case 'B' : return WhiteBishop
	case 'R' : return WhiteRook
	case 'Q' : return WhiteQueen
	case 'K' : return WhiteKing
	case 'p' : return BlackPawn
	case 'n' : return BlackKnight
	case 'b' : return BlackBishop
	case 'r' : return BlackRook
	case 'q' : return BlackQueen
	case 'k' : return BlackKing
	default : return Emptypiece
	}
}