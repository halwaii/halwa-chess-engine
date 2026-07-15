package main

// perft func -> counts all leaf nodes upto certain depths

// base case if depth == 0 return 1
// 1) generate all moves for current board
// 2) loop into all moves
// 3) make move
// 4) check if legal (king is not being attacked)
// 		and add nodes
// 5) unmake move
func Perft(b *board, depth int) uint64{
	// base case : if depth = 0 , then return 1 leave node
	if depth == 0 {
		return 1
	}

	var nodes uint64 = 0

	// 1) generate moves for current board
	var list MoveList
	GenerateAllMoves(b, &list)

	// 2) loop thorugh all moves
	for i:=0; i < len(list.Moves); i++{
		// take each move
		move := list.Moves[i]

		// 3) make move
		MakeMove(b, move)

		// 4) legality check
		// after make move there is toggle in end
		// so to chech whose move was just now we will toggle again
		whomoved := !b.WhiteToMove
		// and check if the king is safe from check or not
		// and then only add it in nodes
		if !isinCheck(*b, whomoved) {
			// we r just saving "legal moves" and removing psuedo legal moves
			nodes += Perft(b, depth - 1)
		}
		// 5) unmake move
		UnMakeMove(b)
	}
	return nodes
}