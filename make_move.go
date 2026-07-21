package main

// make move
func MakeMove (b *board, move Move){

	// 1* remove old board state
	if b.EnPassantSquare != -1{
		b.HashKey ^= ZobristEnpassant[b.EnPassantSquare % 8]
	}
	b.HashKey ^= ZobristTurn
	b.HashKey ^= ZobristCastling[b.CastlingRights]
	// as for pieces we have already done them in add piece 
	// and remove pice ;)
	
	// 1) decode the move
	from := int(GetFrom(move))
	to := int(GetTo(move))
	flag := GetFlag(move)

	// 2) identify piece
	movingPiece := GetPieceAt(b, from)
	capturedPiece := GetPieceAt(b, to)

	// enpassant edge case
	if flag == EpCapture {
		if b.WhiteToMove {
			capturedPiece = BlackPawn
		} else {
			capturedPiece = Whitepawn
		}
	}
	// 4) save the current state and add to our stack
	state := UndoState{
		MoveMade: move,
		Capturedpiece: capturedPiece,
		Enpassantsq: b.EnPassantSquare,
		CastlingRights: b.CastlingRights,
		HalfMoveClock: b.HalfMoveClock,
	}
	b.history = append(b.history, state)

	// core steps for 90% of moves
	// 1) remove the moving piece from current square
	RemovePiece(b, movingPiece, from)

	// 2) remove captured piece (if any)
	if capturedPiece != Emptypiece {
		capturedSquare := to
		// in en passant there is no captured square
		if flag == EpCapture{
			if b.WhiteToMove{
				capturedSquare = to - 8 // new square will be 1 rank below for white
			} else {
				capturedSquare = to + 8 // new square will be 1 rank above for black
			}
		}
		RemovePiece(b, capturedPiece, capturedSquare)
	}
	// 3) set the moving piece to its new  square
	AddPiece(b, movingPiece, to)

	// its time for exceptions :P (special moves)
	// castling and promotion

	// promotion
	if flag >= KnightPromo && flag <= QueenPromoCap{
		// pawn which got added in step 3 should be removed
		RemovePiece(b, movingPiece, to)

		var PromotedPiece int
		if b.WhiteToMove{
			if flag == QueenPromo || flag == QueenPromoCap { PromotedPiece = WhiteQueen
			} else if flag == RookPromo || flag == RookPromoCap { PromotedPiece = WhiteRook
			} else if flag == BishopPromo || flag == BishopPromoCap { PromotedPiece = WhiteBishop
			} else if flag == KnightPromo || flag == KnightPromoCap { PromotedPiece = WhiteKnight}
		} else {
			if flag == QueenPromo || flag == QueenPromoCap { PromotedPiece = BlackQueen
			} else if flag == RookPromo || flag == RookPromoCap { PromotedPiece = BlackRook
			} else if flag == BishopPromo || flag == BishopPromoCap { PromotedPiece = BlackBishop
			} else if flag == KnightPromo || flag == KnightPromoCap { PromotedPiece = BlackKnight}
		}
		// set the promoted piece
		AddPiece(b, PromotedPiece, to)
	}

	// castling
	// king has already made its move so now we have to move rook only
	if flag == KingCastle || flag == QueenCastle{
		if b.WhiteToMove{
			if flag == KingCastle {
				RemovePiece(b, WhiteRook, 7) // clear rook from h1
				AddPiece(b, WhiteRook, 5) // add rook on f1
			} else if flag == QueenCastle {
				RemovePiece(b, WhiteRook, 0) // clear rook from a1
				AddPiece(b, WhiteRook, 3) // add rook on d1
			}
		} else {
			if flag == KingCastle {
				RemovePiece(b, BlackRook, 63) // clear rook from h8
				AddPiece(b, BlackRook, 61) // add rook on f8
			} else if flag == QueenCastle {
				RemovePiece(b, BlackRook, 56) // clear rook from a8
				AddPiece(b, BlackRook, 59) // add rook on d1
			}
		}
	}

	// double pawn push to set enpassnat
	if flag == DoublePawnPush {
		if b.WhiteToMove {
			b.EnPassantSquare = to - 8
		} else {
			b.EnPassantSquare = to + 8
		}
	} else {
		// after every 2nd move en passant square become clear
		b.EnPassantSquare = -1
	}
	// adding 50 move rule
	// a pawn must move or a piece should be captured then reset 0
	if movingPiece == Whitepawn || movingPiece == BlackPawn || capturedPiece != Emptypiece {
		b.HalfMoveClock = 0
	} else {
		b.HalfMoveClock++
	}

	// castling rights update
	// when king moves
	if movingPiece == WhiteKing {
		b.CastlingRights &= ^uint8(3) // here 3 = 0011
	}
	if movingPiece == BlackKing {
		b.CastlingRights &= ^uint8(12) // here 12 = 1100
	}
	// when rook moves or it gets captured 
	// h1 rook -> white king side castling
	if from == 7 || to == 7 { b.CastlingRights &= ^uint8(1)} // here 1 = 0001
	// a1 rook -> white queen side castling
	if from == 0 || to == 0 { b.CastlingRights &= ^uint8(2)} // here 2 = 0010
	// h8 rook -> black king side castling
	if from == 63 || to == 63 { b.CastlingRights &= ^uint8(4)} // here 4 = 0100
	// a8 rook -> black queen side castling
	if from == 56 || to == 56 { b.CastlingRights &= ^uint8(8)} // here 2 = 1000

	// lastly T_T
	// toggle to tell whose move is it next
	b.WhiteToMove = !b.WhiteToMove

	// 2* add new board state in our hash
	if b.EnPassantSquare != -1{
		b.HashKey ^= ZobristEnpassant[b.EnPassantSquare%8]
	}
	b.HashKey ^= ZobristTurn
	b.HashKey ^= ZobristCastling[b.CastlingRights]
}