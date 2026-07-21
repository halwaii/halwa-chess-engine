package main

// 1) remove last state from history
// 2) toggle turn
// 3) restore state : castling, en passant, 50 move rule
// 4) move piece back
// 5) restore captured piece
// 6) reverse castling
func UnMakeMove(b *board){

	// 1* remove current board state
	if b.EnPassantSquare != -1{
		b.HashKey ^= ZobristEnpassant[b.EnPassantSquare%8]
	}
	b.HashKey ^= ZobristTurn
	b.HashKey ^= ZobristCastling[b.CastlingRights]

	// 1) save the last position from history and pop it
	lastIndex := len(b.history) - 1
	state := b.history[lastIndex]
	// poping is done by [:]
	b.history = b.history[:lastIndex]

	// 2) toggle turn
	b.WhiteToMove = !b.WhiteToMove

	// 3) resotre state to old state
	b.EnPassantSquare = state.Enpassantsq
	b.CastlingRights = state.CastlingRights
	b.HalfMoveClock = state.HalfMoveClock

	// 4) move piece back

	// extract move details
	move := state.MoveMade
	from := int(GetFrom(move))
	to := int(GetTo(move))
	flag := GetFlag(move)

	// now the moving piece is at "to"
	movedPiece := GetPieceAt(b, to)

	// reverse promotion
	if flag >= KnightPromo && flag <= QueenPromoCap {
		// then remove promoted piece 
		RemovePiece(b, movedPiece, to)
		// promoted piece was a pawn before
		if b.WhiteToMove{
			movedPiece = Whitepawn
		} else {
			movedPiece = BlackPawn
		}
	} else {
		// for normal cases we just remove piece
		RemovePiece(b, movedPiece, to)
	}
	// add the piece to its original position
	AddPiece(b, movedPiece, from)

	// 5) restore captured piece -> Zombieeeee
	if state.Capturedpiece != Emptypiece {
		capturedSquare := to

		// enpassant speical
		if flag == EpCapture {
			if b.WhiteToMove {
				capturedSquare = to - 8 // black pawn was 1 row behind
			} else {
				capturedSquare = to + 8 // white pawn was 1 row above
			}
		}
		// bug bug bug
		// if b.WhiteToMove {
		// 	capturedSquare = to - 8 // black pawn was 1 row behind
		// } else {
		// 	capturedSquare = to + 8 // white pawn was 1 row above
		// }
		// add the piece where it was captured
		AddPiece(b, state.Capturedpiece, capturedSquare)
	}
	
	// 6) reverse castling (very tuff XD)
	// just remove piece from castled position to
	// original position
	if flag == KingCastle || flag == QueenCastle {
		if b.WhiteToMove{
			if flag == KingCastle {
				RemovePiece(b, WhiteRook, 5)
				AddPiece(b, WhiteRook, 7)
			} else if flag == QueenCastle {
				RemovePiece(b, WhiteRook, 3)
				AddPiece(b, WhiteRook, 0)
			}
		} else {
			if flag == KingCastle {
				RemovePiece(b, BlackRook, 61)
				AddPiece(b, BlackRook, 63)
			} else if flag == QueenCastle {
				RemovePiece(b, BlackRook, 59)
				AddPiece(b, BlackRook, 56)
			}
		}
	}

	// 2* add new board state in our hash
	if b.EnPassantSquare != -1{
		b.HashKey ^= ZobristEnpassant[b.EnPassantSquare%8]
	}
	b.HashKey ^= ZobristTurn
	b.HashKey ^= ZobristCastling[b.CastlingRights]
}