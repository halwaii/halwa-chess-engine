package main

// how much max depth will we search in tree
const maxPly = 64

// two killer moves (quiet, non hash) per ply 
var killerMoves [maxPly][2]Move

// history heuristic 
var historyTable [13][64]int

// function to store killer moves
func StoreKiller(move Move, ply int){
	// check boundaries
	if ply<0 || ply >= maxPly{
		return
	}
	if move == killerMoves[ply][0]{
		return // already top killer move
	}
	// if found better move then store it as killer move
	killerMoves[ply][1] = killerMoves[ply][0]
	killerMoves[ply][0] = move
}

func storeHistory(piece int, to int, depth int){
	// boundaries
	if piece<0 || piece>12 || to <0 || to>63{
		return
	}
	// the move with more depth have way more score than the move with low depth
	historyTable[piece][to] += depth*depth
}

// clear arrays function
func ClearKillerHistory(){
	killerMoves= [maxPly][2]Move{}
	historyTable=[13][64]int{}
}