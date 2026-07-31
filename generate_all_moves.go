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
		WhitePawnmoves(*b, list, false)
		allLegalKnightmoves(*b, true, list, false)
		allLegalBishopmoves(*b, true, list, false)
		allLegalRookmoves(*b, true, list, false)
		allLegalQueenmoves(*b, true, list, false)
		allLegalKingmoves(*b, true, list, false)
	} else {
		BlackPawnmoves(*b, list, false)
		allLegalKnightmoves(*b, false, list, false)
		allLegalBishopmoves(*b, false, list, false)
		allLegalRookmoves(*b, false, list, false)
		allLegalQueenmoves(*b, false, list, false)
		allLegalKingmoves(*b, false, list, false)
	}
}

func GenerateCapturesOnly(b *board, list *MoveList){
	// same as generate all moves 
	// just add true for captures only 
	// before generating new moves the list should be empty
	list.count = 0
	list.Moves = nil

	if b.WhiteToMove{
		// call all legal moves generation function
		WhitePawnmoves(*b, list, true)
		allLegalKnightmoves(*b, true, list, true)
		allLegalBishopmoves(*b, true, list, true)
		allLegalRookmoves(*b, true, list, true)
		allLegalQueenmoves(*b, true, list, true)
		allLegalKingmoves(*b, true, list, true)
	} else {
		BlackPawnmoves(*b, list, true)
		allLegalKnightmoves(*b, false, list, true)
		allLegalBishopmoves(*b, false, list, true)
		allLegalRookmoves(*b, false, list, true)
		allLegalQueenmoves(*b, false, list, true)
		allLegalKingmoves(*b, false, list, true)
	}
}