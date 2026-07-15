package main

// this is master funciton of all pieces moves
// this function stores all the possible moves of current player
// in our moves list wrt board state
func GenerateAllMoves(b *board, list *MoveList){
	// before generating new moves the list should be empty
	list.count = 0
	list.Moves = nil

	if b.WhiteToMove{
		// call all legal moves generation function
		WhitePawnmoves(*b, list)
		allLegalKnightmoves(*b, true, list)
		allLegalBishopmoves(*b, true, list)
		allLegalRookmoves(*b, true, list)
		allLegalQueenmoves(*b, true, list)
		allLegalKingmoves(*b, true, list)
	} else {
		BlackPawnmoves(*b, list)
		allLegalKnightmoves(*b, false, list)
		allLegalBishopmoves(*b, false, list)
		allLegalRookmoves(*b, false, list)
		allLegalQueenmoves(*b, false, list)
		allLegalKingmoves(*b, false, list)
	}
}