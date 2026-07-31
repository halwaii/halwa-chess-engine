package main

import (
	"sort"
)

// maximum and minimum score
const(
	infinity = 50000
	neginfinity = -50000
)

// alpha beta search function (negamax)
// alpha -> my best case
// beta -> opponents best case
func Search(b *board, depth int, alpha int, beta int) int{
	// if depth = 0 then evaluate the score of current board
	if depth == 0{
		return QuiescenceSearch(b, alpha, beta)
	}
	// maximum starts from negative infinity
	bestscore := neginfinity

	// generate all legal moves in current board
	var list MoveList
	GenerateAllMoves(b, &list)

	// move ordering should be before doing dfs
	// creating a new slice to store moves and its score
	type ScoredMove struct{
		move Move
		score int
	}
	// allocating memory and calculating scores
	// we use "make" for it
	scoredMoves := make([]ScoredMove, len(list.Moves))
	for i:=0;i<len(list.Moves);i++{
		scoredMoves[i] = ScoredMove{
			move: list.Moves[i],
			score: ScoreMoves(b, list.Moves[i]),
		}
	}

	// sort the slice in descending order
	sort.Slice(scoredMoves, func(i,j int)bool{
		return scoredMoves[i].score>scoredMoves[j].score
	})

	// count legal moves
	LegalMovesCount := 0
	// loop through all moves
	for i:=0;i<len(list.Moves);i++{
		move := scoredMoves[i].move // update here
		// make move
		MakeMove(b, move)

		// check for legality
		// if after the move the king is in check or not
		// if yes then continue
		if isinCheck(*b, !b.WhiteToMove){
			// if yes them immediate unmakemove
			UnMakeMove(b)
			continue
		}
		LegalMovesCount++
		// recurse with NEGAmax
		// negamax : depth--, flip alpha beta and negate too, and negate score
		// max(a, b) = -min(-a, -b)
		score := -Search(b, depth-1, -beta, -alpha)

		// unmake move
		UnMakeMove(b)

		// update the new score if it is better than last
		if score > bestscore{
			bestscore = score
		}

		// update alpha
		// alpha acts like max in minimax
		if score > alpha{
			alpha = score
		}

		// alpha beta pruning
		// score is already max so no need to check further
		if alpha >= beta{
			break // 
		}
	}
	// logic for checkmate and stalemate
	if LegalMovesCount==0{
		// if there is no legal moves and king is in check -> CHECKMATE
		if isinCheck(*b, b.WhiteToMove){
			return neginfinity - (100-depth) // to find mate faster
		}
		// if not in check but there are no legal moves -> STALEMATE
		return  0
	}
	return bestscore
}