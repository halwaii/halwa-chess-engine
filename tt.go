package main

// using zobrist we make unique ID of a board state
// and now we will save this chache memory 
// we call this transposition table

// we need 3 things
// 1) hashkey : unique ID of board
// 2) depth
// 3) nodes 

// define TT
type TTEntry struct{
	Depth int
	Nodes uint64
}

// declare a global hashmap whose key will be 64 bit zobrist hash
// and value will be TTEntry struct

var TranspositionTable map[uint64]TTEntry