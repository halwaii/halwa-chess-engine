package main

const developmentCP = -30

// penalties for staying in 1st rank or 8th rank during middle game
func evalDevelopment(b *board) int{
	score := 0

	if (b.WhiteKnight & (uint64(1)<<uint64(1))) !=0 {
		score += developmentCP
	}
	if (b.WhiteKnight & (uint64(1)<<uint64(6))) !=0 {
		score += developmentCP
	}
	if (b.WhiteBishop & (uint64(1)<<uint64(2))) !=0 {
		score += developmentCP
	}
	if (b.WhiteBishop & (uint64(1)<<uint64(5))) !=0 {
		score += developmentCP
	}

	if (b.BlackKnight & (uint64(1)<<uint64(57))) !=0 {
		score -= developmentCP
	}
	if (b.BlackKnight & (uint64(1)<<uint64(62))) !=0 {
		score -= developmentCP
	}
	if (b.BlackBishop & (uint64(1)<<uint64(58))) !=0 {
		score -= developmentCP
	}
	if (b.BlackBishop & (uint64(1)<<uint64(61))) !=0 {
		score -= developmentCP
	}

	return  score
}