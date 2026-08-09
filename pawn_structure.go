package main

import "math/bits"

const (
	doubledPawnMG   = -10 // two or more same side pawns on same file
	doubledPawnEG   = -20
	isolatedPawnMG  = -12 // single pawn
	isolatedPawnEG  = -18
	supportedPawnMG = 8 // pawn defended by other pawn
	supportedPawnEG = 12
	connectedPawnMG = 6 // pawn on adjecent file (col)
	connectedPawnEG = 10
)

// func to make file mask of given file
// init() function runs even before main
var FileMasks[8] uint64
func init(){
	for c:=0;c<8;c++{
		// chose every column
		var mask uint64
		for r:=0;r<8;r++{
			mask |= uint64(1)<<uint64(r*8+c)
		}
		FileMasks[c]=mask
	}
}
// indexed by rank
var passedPawnBonusMG = [8]int{0, 5, 10, 15, 25, 40, 60, 0}
var passedPawnBonusEG = [8]int{0, 10, 25, 40, 65, 100, 150, 0}

func evalPawn(ownPawns uint64, enemyPawns uint64, isWhite bool) (int, int) {
	mg, eg := 0, 0
	// make local copy, so that original doesnt get modified
	pawns := ownPawns
	// loop till no pawns left
	for pawns != 0 {
		square := bits.TrailingZeros64(pawns)
		col := square%8
		row := square/8

		// doubled (same side pawns on same files)
		// checks if pawn is +nt on same file
		if bits.OnesCount64(ownPawns & FileMasks[col]) > 1{
			mg += doubledPawnMG
			eg += doubledPawnEG
		}

		// isolated
		var neighborMask uint64 // all zero
		if col>0{neighborMask |= FileMasks[col - 1]}
		if col<7{neighborMask |= FileMasks[col + 1]}
		// if no pawn +nt on neighboring col
		isolated := (neighborMask & ownPawns)==0
		if isolated{
			mg += isolatedPawnMG
			eg += isolatedPawnEG
		}

		// supported -> 1 row behind and present in adjacent file diagonally
		behindRow := row -1
		if !isWhite{ // for black
			behindRow = row+1
		}
		if behindRow >=0 && behindRow<8{
			if col>0 && (ownPawns & (uint64(1)<<uint64(behindRow*8 + col-1))) !=0{
				mg += supportedPawnMG
				eg += supportedPawnEG
			} else if col<7 && (ownPawns & (uint64(1)<<uint64(behindRow*8 + col+1))) !=0{
				mg += supportedPawnMG
				eg += supportedPawnEG
			}
		}

		// all connected pawns are supported but all supported ar not connected
		if !isolated{
			if col>0 && (ownPawns & (uint64(1)<<uint64(row*8 + col-1))) !=0{
				mg += supportedPawnMG
				eg += supportedPawnEG
			} else if col>0 && (ownPawns & (uint64(1)<<uint64(row*8 + col+1))) !=0{
				mg += supportedPawnMG
				eg += supportedPawnEG
			}
		}
		// passed pawn
		// checks if there is no enemy pawn on same file and adjacent files
		// creates a mask to check enemy pawns
		var frontSpan uint64
		step := 1
		if !isWhite{
			step = -1
		}
		// check further ranks from curr rank
		for r:=row+step; r>=0 && r<8; r += step {
			frontSpan |= uint64(1)<<uint64(r*8+col) 
			if col>0 {
				frontSpan |= uint64(1)<<uint64(r*8+col-1)
			}
			if col<7 {
				frontSpan |= uint64(1)<<uint64(r*8+col+1)
			}
		}
		// & with enemyPawns and frontSpan
		if (enemyPawns & frontSpan) == 0{
			// it is a passed pawn so give bonus
			relRow := row
			if !isWhite{
				relRow = 7-row
			}
			mg += passedPawnBonusMG[relRow]
			eg += passedPawnBonusEG[relRow]
		}
		// decrese by 1
		pawns &= pawns-1
	}
	return mg, eg
}
func EvaluatePawnStructure(b *board) (int, int) {
	wMG, wEG := evalPawn(b.WhitePawns, b.BlackPawns, true)
	bMG, bEG := evalPawn(b.BlackPawns, b.WhitePawns, false)
	return wMG - bMG, wEG - bEG
}