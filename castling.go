package main

const (
	castlingBonus = 40
	LostCastling = -50 // lost rights without castling -> penalty
)

func evalCastling(b *board) int {
	score := 0

	if (b.WhiteKing & (uint64(1)<<6)) != 0 && (b.WhiteRook & (uint64(1)<<5)) != 0 {
		score += castlingBonus
	}
	if (b.WhiteKing & (uint64(1)<<2)) != 0 && (b.WhiteRook & (uint64(1)<<3)) != 0 {
		score += castlingBonus
	}
	if (b.BlackKing & (uint64(1)<<62)) != 0 && (b.BlackRook & (uint64(1)<<61)) != 0 {
		score -= castlingBonus
	}
	if (b.BlackKing & (uint64(1)<<58)) != 0 && (b.BlackRook & (uint64(1)<<59)) != 0 {
		score -= castlingBonus
	}
	return score
}