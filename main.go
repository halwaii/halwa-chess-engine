package main

func main(){
	// var myMove Move = EncodeMove(12, 28, 0)
	// fmt.Println("pawn moves from e2(12) to e4(28) : \n")
	// fmt.Println("Encoded Move Value (Integer):", myMove)
	// fmt.Println("From Square:", myMove.GetFrom())
	// fmt.Println("To Square:", myMove.GetTo())
	// fmt.Println("Move Flag:", myMove.GetFlag())
	var b board
	b.WhitePawns = 0x000000000000FF00
	b.WhiteKing = 0x0000000000000010
	b.WhiteQueen = 0x0000000000000008
	b.WhiteBishop = 0x0000000000000024 
	b.WhiteKnight = 0x0000000000000042
	b.WhiteRook = 0x0000000000000081
	
	b.BlackPawns = 0x00FF000000000000
	b.BlackKing = 0x1000000000000000
	b.BlackRook = 0x8100000000000000
	b.BlackBishop = 0x2400000000000000
	b.BlackKnight = 0x4200000000000000
	b.BlackQueen = 0x0800000000000000

	Printboard(b)

	// b.WhiteQueen = uint64(1) << 28
	// b.BlackPawns = uint64(1) << 46
	// fmt.Println("\nwhite queen legal moves(d4) and a blackpiece at g6 : \n")
	// queenmoves := allLegalQueenmoves(b, true)
	// PrintBitboard(queenmoves)
	// rookmoves := allLegalBlackRookmoves(b)
	// fmt.Println("\nblack Rook Legal Moves when it is at d4 and h8 and there is black piece(f4) and a white piece(d6):\n")
	// PrintBitboard(rookmoves)
	// whiteknightmoves := allLegalWhiteKnightmoves(b)
	// blackknightmoves := allLegalBlackKnightmoves(b)
	// fmt.Println("White Knight Legal Moves Matrix:")
	// PrintBitboard(whiteknightmoves)

	// fmt.Println("Black Knight Legal Moves Matrix:")
	// PrintBitboard(blackknightmoves)
}