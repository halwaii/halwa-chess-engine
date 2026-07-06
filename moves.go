package main

// move pawn
func Whitepawnpush(b board) uint64{
	// single push
	Singlepush := (b.WhitePawns << 8) & ^Occupiedsquares(b)

	// double push
	// and row 3 and row 4 should be empty
	row3mask := uint64(0x000000000000FF0000)
	Doublepush := ((Singlepush & row3mask)<<8) & ^Occupiedsquares(b)

	// capture
	// capture can be left (<< 7) but it should not be on A-lane and it can only capture black piece
	Leftcapture := ((b.WhitePawns & notA_lane) << 7) & blackpieces(b)

	// capture can be right (<< 9) but it should not be on H-lane and there should be black piece
	Rightcapture := ((b.WhitePawns & notH_lane) << 9) & blackpieces(b)

	// dynamic En passant 
	var LeftEnpassant uint64 = 0
	var RightEnpassant uint64 = 0

	if b.EnPassantSquare != -1 {
		LeftEnpassant = ((b.WhitePawns & notA_lane) << 7) & (uint64(1)<<uint64(b.EnPassantSquare))
		RightEnpassant = ((b.WhitePawns & notH_lane) << 9) & (uint64(1)<<uint64(b.EnPassantSquare))
	}

	return Singlepush | Doublepush | Leftcapture | Rightcapture | LeftEnpassant | RightEnpassant
}

func Blackpawnpush(b board) uint64{
	// single push
	Singlepush := (b.BlackPawns >> 8) & ^Occupiedsquares(b)

	// double push
	// row 6 and 5 should be empty
	mask6row := uint64(0x0000FF0000000000)
	Doublepush := ((Singlepush & mask6row) >> 8) & ^Occupiedsquares(b)

	// capture
	// capture00 can be left (>> 9) and not on A-lane
	Leftcapture := ((b.BlackPawns & notA_lane) >> 9) & whitepieces(b)
	// capture can be right (>> 7) and not on H-lane
	Rightcapture := ((b.BlackPawns & notH_lane) >> 7) & whitepieces(b)

	// dynamic en passant
	var LeftEnpassant uint64 = 0
	var RightEnpassant uint64 = 0
	if b.EnPassantSquare != -1 {
		LeftEnpassant = ((b.BlackPawns & notA_lane) >> 7) & (uint64(1)<<uint64(b.EnPassantSquare))
		RightEnpassant = ((b.BlackPawns & notH_lane) >> 9) & (uint64(1)<<uint64(b.EnPassantSquare))
	}

	return Singlepush | Doublepush | Leftcapture | Rightcapture | LeftEnpassant | RightEnpassant
}


// all possible knight moves
func Knightmoves(square int) uint64{
	// create moves to store all possible moves of knight
	var moves uint64 = 0
	// first find row and column of our knight
	row := square / 8
	col := square % 8

	// all 8 - L shaped possible offsets of knight
	rowOffsets := []int{2,2,1,1,-1,-1,-2,-2}
	colOffsets := []int{1,-1,2,-2,2,-2,1,-1}

	for i:=0 ; i<8; i++ {
		// finding rows and columns of possible moves
		r := row + rowOffsets[i]
		c := col + colOffsets[i]
		
		// edge cases
		// possible moves should not go beyond
		if r>=0 && r<8 && c>=0 && c<8 {
			PossibleSquare := r*8 + c
			moves |= ((uint64(1) << uint64(PossibleSquare)))
		}
	}
	return moves
}
// all Legal white Knight moves
func LegalKnightmoves(b board, square int, isWhite bool) uint64{
	rawmoves := Knightmoves(square)
	// so that white knight cannnot take white pieces
	if isWhite {
		return rawmoves & ^whitepieces(b)
	} else {
		return rawmoves & ^blackpieces(b)
	}
}
// all Legal Black Knight moves
// func LegalBlackKnightmoves(b board, square int) uint64{
// 	rawmoves := Knightmoves(square)
// 	// so that black knight cannnot take black pieces
// 	return rawmoves & ^blackpieces(b)
// }
// we will find the square on which our knight is actually standing
func allLegalKnightmoves(b board, isWhite bool) uint64{
	var allKnightmoves uint64 = 0
	for square := 0 ; square < 64 ; square++ {
		var currKnight uint64 = 0
		if isWhite {
			currKnight = b.WhiteKnight
		} else {
			currKnight = b.BlackKnight
		}
		if (currKnight & (uint64(1) << uint64(square))) != 0 {
			allKnightmoves |= LegalKnightmoves(b, square, isWhite)
		}
	}
	return allKnightmoves
}
// func allLegalBlackKnightmoves(b board) uint64{
// 	var allBlackKnightmoves uint64 = 0
// 	for square := 0 ; square < 64 ; square++ {
// 		if (b.BlackKnight & (uint64(1) << uint64(square))) != 0 {
// 			allBlackKnightmoves |= LegalBlackKnightmoves(b, square)
// 		}
// 	}
// 	return allBlackKnightmoves
// }

// all possible sliding moves
func Slidingmoves(square int, b board, rowOffsets, colOffsets []int) uint64{
	var moves uint64 = 0
	// find row and col of current rook/bishop/queen
	row := square / 8
	col := square % 8

	// ray casting in every direction
	// here instead of 4 we will do len(rowOffsets) 
	// cause queen can move in 8 directions
	for i :=0;i< len(rowOffsets) ;i++ {
		r := row + rowOffsets[i]
		c := col + colOffsets[i]
		// going on till end of board
		for r>=0 && r<8 && c>=0 && c<8 {
			PossibleSquare := r*8 + c
			mask := (uint64(1) << uint64(PossibleSquare))
			moves |= mask
			// break considtion if a square is occupied
			if (Occupiedsquares(b) & mask) != 0 {
				break
			}
			// increment/decrement
			r += rowOffsets[i]
			c += colOffsets[i]
		}
	}
	return moves
}
// instead of writing different functions for black and white we wrote in single funciton
// get all Legal rook moves
func LegalRookmoves(b board, square int, isWhite bool) uint64{
	// rook can move in EAST WEST NORTH SOUTH directions
	rowOffsets := []int{1,-1,0,0}
	colOffsets := []int{0,0,1,-1}
	rawmoves := Slidingmoves(square, b, rowOffsets, colOffsets)

	if isWhite {
		return rawmoves & ^whitepieces(b)
	} else {
		return rawmoves & ^blackpieces(b)
	}
}
// find squares where rooks are present
func allLegalRookmoves(b board, isWhite bool) uint64{
	var allRookmoves uint64 = 0
	for square:=0 ; square<64; square++{
		var currentRook uint64 = 0
		if isWhite {
			currentRook = b.WhiteRook
		} else {
			currentRook = b.BlackRook
		}
		if (currentRook & (uint64(1)<<uint64(square))) != 0{
			allRookmoves |= LegalRookmoves(b, square, isWhite)
		}
	}
	return allRookmoves
}
// same logic is for bishops
func LegalBishopmoves(b board, square int, isWhite bool) uint64{
	// bishop moves in NE , NW , SE , SW
	rowOffsets := []int{1,1,-1,-1}
	colOffsets := []int{1,-1,1,-1}
	
	rawmoves := Slidingmoves(square, b, rowOffsets, colOffsets)

	if isWhite {
		return rawmoves & ^whitepieces(b)
	} else {
		return rawmoves & ^blackpieces(b)
	}
}
func allLegalBishopmoves(b board, isWhite bool) uint64{
	var allBishopmoves uint64 = 0
	for square := 0 ; square < 64 ; square++ {
		var currentBishop uint64 = 0
		if isWhite {
			currentBishop = b.WhiteBishop
		} else {
			currentBishop = b.BlackBishop
		}
		if (currentBishop & (uint64(1)<<uint64(square))) != 0 {
			allBishopmoves |= LegalBishopmoves(b, square, isWhite)
		}
	}
	return allBishopmoves
}
// queen = bishop + rook
// it can move in 8 directions
func allLegalQueenmoves(b board, isWhite bool) uint64{
	var allQueenmoves uint64 = 0
	for square :=0 ; square<64; square++ {
		var currentQueen uint64 = 0
		if isWhite {
			currentQueen = b.WhiteQueen
		} else {
			currentQueen = b.BlackQueen
		}
		if (currentQueen & (uint64(1)<<uint64(square))) != 0{
			allQueenmoves |= LegalBishopmoves(b, square, isWhite) | LegalRookmoves(b, square, isWhite)
		}
	}
	return allQueenmoves
}
// all possisble king moves
func Kingmoves(square int) uint64{
	var moves uint64 = 0
	// find row and col of king's position
	row := square / 8
	col := square % 8
	// offsets for 8 directions
	rowOffsets := []int{1,1,1,0,0,-1,-1,-1}
	colOffsets := []int{-1,0,1,1,-1,-1,0,1}

	for i:=0; i<8; i++{
		// next row and cols
		r := row + rowOffsets[i]
		c := col + colOffsets[i]

		if r>=0 && r<8 && c>=0 && c<8 {
			// target square
			PossibleSquare := r*8 + c
			moves |= (uint64(1)<<uint64(PossibleSquare))
		}
	}
	return moves
}
func LegalKingmoves(b board, square int, isWhite bool) uint64{
	rawmoves := Kingmoves(square)
	// check for castling
	var finalmoves uint64 = 0
	if isWhite {
		finalmoves = rawmoves & ^whitepieces(b)
		// white castle checking . king must be on e1
		if square == 4 {
			//king side castle from e1 -> g1
			// all the space should be empty
			if (Occupiedsquares(b) & whiteKingsidemask) == 0 {
				// and should not be attacked by black piece
				// e1 , f1, g1 should not be attacked by black
				if !IsSquareAttacked(4,false,b) && !IsSquareAttacked(5,false,b) && !IsSquareAttacked(6,false,b) {
					finalmoves |= uint64(1)<<6
				}
			}
			// queen side castle from e1 to c1
			if (Occupiedsquares(b) & whiteQueensidemask) == 0 {
				if !IsSquareAttacked(2,false,b) && !IsSquareAttacked(3,false,b) && !IsSquareAttacked(4,false,b){
					finalmoves |= uint64(1)<<2
				}
			}
		}
	} else {
		finalmoves = rawmoves & ^blackpieces(b)
		// black castling check . king must be on e8
		if square == 60 {
			// king side from e8 -> g8
			if (Occupiedsquares(b) & blackKingsidemask) == 0 {
				// e8, f8, g8 should not be attacked by white
				if !IsSquareAttacked(60,true,b) && !IsSquareAttacked(61,true,b) && !IsSquareAttacked(62,true,b) {
					finalmoves |= uint64(1)<<62
				}
			}
			// queen side from e8 -> c8
			if (Occupiedsquares(b) & blackQueensidemask) == 0 {
				if !IsSquareAttacked(58,true,b) && !IsSquareAttacked(59,true,b) && !IsSquareAttacked(60,true,b) {
					finalmoves |= uint64(1)<<58
				}
			}
		}
	}
	return finalmoves
}
func allLegalKingmoves(b board, isWhite bool) uint64{
	var allKingmoves uint64 = 0
	for square:=0; square<64; square++ {
		var currKing uint64 = 0;
		if isWhite{
			currKing = b.WhiteKing
		} else {
			currKing = b.BlackKing
		}
		if (currKing & (uint64(1)<<uint64(square))) != 0{
			allKingmoves |= LegalKingmoves(b, square, isWhite)
		}
	}
	return allKingmoves
}
func isinCheck(b board, isWhite bool) bool{
	var kingSquare int
	var kingBitboard uint64 = 0
	if isWhite {
		kingBitboard = b.WhiteKing
	} else {
		kingBitboard = b.BlackKing
	}
	// we find find square of that particular king
	for square :=0; square<64; square++ {
		if (kingBitboard & (uint64(1)<<uint64(square))) != 0 {
			kingSquare = square
			break
		}
	}
	// if white is in check then it should be attacked by black piece
	// so !isWhite kiya
	return IsSquareAttacked(kingSquare, !isWhite, b)
}
// func LegalBlackRookmoves(b board, square int) uint64{
// 	rowOffsets := []int{1,-1,0,0}
// 	colOffsets := []int{0,0,1,-1}
// 	rawmoves := Slidingmoves(square, b, rowOffsets, colOffsets)
// 	// filter so that they can't take their friendly piece
// 	return rawmoves & ^blackpieces(b)
// }

// func allLegalBlackRookmoves(b board) uint64{
// 	var allRookmoves uint64 = 0
// 	for square:=0 ; square<64 ; square++ {
// 		if (b.BlackRook & (uint64(1) << uint64(square))) != 0{
// 			allRookmoves |= LegalBlackRookmoves(b, square)
// 		}
// 	}
// 	return allRookmoves
// }

// is our target square under attacked ?
// we will be using revese logic . 
// we will check for all moves of other pieces which will fall on that square
func IsSquareAttacked(square int, isWhite bool, b board) bool{
	// check attacks form pawn
	if isWhite {
		// row should be more than 0
		if square/8 > 0 {
			if square%8 > 0 && (b.WhitePawns & (uint64(1)<<uint64(square-9))) != 0 { return true}
			if square%8 < 7 && (b.WhitePawns & (uint64(1)<<uint64(square-7))) != 0 { return true}
		}
	} else {
		// row should be less than 7
		if square/8 < 7 {
			// check for A_lane and H_lane
			if square%8 > 0 && (b.BlackPawns & (uint64(1)<<uint64(square+7))) != 0 { return true}
			if square%8 < 7 && (b.BlackPawns & (uint64(1)<<uint64(square+9))) != 0 { return true}
		}
	}
	// check attacks form knight
	// we will find knight moves on that particular square and then 
	// check if knight is present on those square
	if isWhite {
		if (Knightmoves(square) & b.WhiteKnight) != 0 { return true}
	} else {
		if (Knightmoves(square) & b.BlackKnight) != 0 { return true}
	}
	// check attacks form rook / queen (straight line)
	// cast rays for dummy rook from given square
	rookAttacks := Slidingmoves(square, b, []int{1,-1,0,0}, []int{0,0,1,-1})
	if isWhite {
		if (rookAttacks & (b.WhiteRook | b.WhiteQueen)) != 0 { return true}
	} else {
		if (rookAttacks & (b.BlackRook | b.BlackQueen)) != 0 {return true}
	}
	// check attacks from bishop / queen (diagonal)
	bishopAttacks := Slidingmoves(square, b, []int{1,1,-1,-1}, []int{1,-1,1,-1})
	if isWhite {
		if (bishopAttacks & (b.WhiteBishop | b.WhiteQueen)) != 0 {return true}
	} else {
		if (bishopAttacks & (b.BlackBishop | b.BlackQueen)) != 0 {return true}
	}
	// check attacks from king
	if isWhite {
		if (Kingmoves(square) & b.WhiteKing) != 0 {return true}
	} else {
		if (Kingmoves(square) & b.BlackKing) != 0 {return true}
	}
	return false
}