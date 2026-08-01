package main

// using zobrist we make unique ID of a board state
// and now we will save this chache memory 
// we call this transposition table

// we need 3 things
// 1) hashkey : unique ID of board
// 2) depth
// 3) nodes 

// TT flags to know what kind of score we stored
const (
	Exactflag=0 // perfect score
	Alphaflag = 1 // upper bound (engine failed to bring good score)
	Betaflag = 2 // lower bound (engine got beta cut off)
)
// define TT
type TTEntry struct{
	Depth int
	Score int
	Flag int
	Nodes uint64
	BestMove Move
}

// declare a global hashmap whose key will be 64 bit zobrist hash
// and value will be TTEntry struct

var TranspositionTable map[uint64]TTEntry