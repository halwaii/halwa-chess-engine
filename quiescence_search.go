package main

import "sort"
// keeps searching unless quiet/peace is achieved on board even if depth has become 0
func QuiescenceSearch(b *board, alpha int, beta int)int{
	// 1) stand-pat score (lazy evaluation)
	// check normal score ie without capture
	standPat := Evaluate(b)

	// get out of function
	// prune
	if standPat >= beta{
		return beta
	}
	// same as our main Search
	if standPat > alpha{
		alpha = standPat
	}

	// 2) generate captures only 
	// we only need captures to keep the search running
	var list MoveList
	GenerateCapturesOnly(b, &list)

	// move ordering should be before doing dfs
	// creating a new slice to store moves and its score
	type ScoredMove struct{
		move Move
		score int
	}
	// can also add move ordering to make it faster
	// same as in search
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

	// loop through captures
	for i:=0;i<len(scoredMoves);i++{
		move := scoredMoves[i].move
		MakeMove(b, move)
		// legality check
		if isinCheck(*b, !b.WhiteToMove){
			UnMakeMove(b)
			continue
		}
		// recursion with negamax logic
		score := -QuiescenceSearch(b, -beta, -alpha)

		UnMakeMove(b)

		// alpha beta pruning
		if score>=beta{
			return beta
		}
		if score>alpha{
			alpha = score
		}
	}
	return alpha
}