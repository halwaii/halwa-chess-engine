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

	// alpha keeps changing, so we store it in TT
	// 1) store original alpha to decide flag in TT
	ogAlpha := alpha

	// 1* create variable of to hash beste move
	var hashMove Move
	// 2) TT read
	if entry, found := TranspositionTable[b.HashKey];found{

		// 2* save old best move from cache
		hashMove = entry.BestMove

		// use cache only when old depth is bigger than curr 
		if entry.Depth >= depth{
			if entry.Flag == Exactflag{
				return entry.Score // return exact score
			}
			if entry.Flag == Alphaflag && entry.Score <= alpha{
				return alpha // upper bound prune
			}
			if entry.Flag == Betaflag && entry.Flag >= beta{
				return beta // lower bound prune
			}
		}
	}
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

		// normal evaluation
		moveScore := ScoreMoves(b, list.Moves[i])

		// if this move was caches best move , then give it godly score
		if list.Moves[i] == hashMove{
			// so that it becomes top move when sorted 
			// and alpha beta pruning will do its work
			moveScore = 10000000
		}
		scoredMoves[i] = ScoredMove{
			move: list.Moves[i],
			score: moveScore,
		}
	}

	// sort the slice in descending order
	sort.Slice(scoredMoves, func(i,j int)bool{
		return scoredMoves[i].score>scoredMoves[j].score
	})

	// count legal moves
	LegalMovesCount := 0

	// save the new BestMove
	// to track curr best move
	var currBestMove Move

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
			currBestMove = move
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
	// 3) TT write
	flag := Exactflag // let exact flag be default
	if bestscore<= ogAlpha{
		flag = Alphaflag 
	} else if bestscore >= beta{
		flag = Betaflag
	}
	TranspositionTable[b.HashKey] = TTEntry{
		Depth: depth,
		Score: bestscore,
		Flag: flag,
		BestMove: currBestMove,
	}
	return bestscore
}