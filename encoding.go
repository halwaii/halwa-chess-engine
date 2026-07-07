package main

// move encoding 
type Move uint16
// 2^6 = 64
// we need 6 bits to define each square
// 0 to 5 bits - from
// 6 to 11 bits - to
// 12 to 15 bits - flag

// 4-bit flags for defining move properties
const(
	// promotion capture special_1 special_0
	// 	   0 		0		0		  0				
	QuietMove 		uint16 = 0 	// 0000
	DoublePawnPush 	uint16 = 1 	// 0001
	KingCastle 		uint16 = 2 	// 0010
	QueenCastle 	uint16 = 3 	// 0011
	Capture 		uint16 = 4 	// 0100
	EpCapture 		uint16 = 5 	// 0101
	// there are only 2 captures -> normal and enpassnt
	// there is no 3rd special capture in chess
	// so these moves doesn't exist and are left as it is
	// 6 -> 0110 
	// 7 -> 0111
	KnightPromo 	uint16 = 8	// 1000
	BishopPromo 	uint16 = 9	// 1001
	RookPromo 		uint16 = 10 // 1010
	QueenPromo 		uint16 = 11 // 1011
	KnightPromoCap 	uint16 = 12 // 1100
	BishopPromoCap 	uint16 = 13 // 1101
	RookPromoCap 	uint16 = 14 // 1110
	QueenPromoCap 	uint16 = 15 // 1111
)
// func to encode move
func EncodeMove(from uint16, to uint16, flag uint16) Move{
	// bitwise OR to make single 16 bit number
	return Move(from | (to << 6) | (flag << 12))
}
// move decode
func GetFrom(m Move) uint16{
	return uint16(m) & 0x3F // 0x3F is 63
}
func GetTo(m Move) uint16{
	return (uint16(m) >> 6) & 0x3F 
}
func GetFlag(m Move) uint16{
	return (uint16(m) >> 12) & 0x3F
}
// checks for 3rd bit (from rigth to left)
// if 3rd bit is 1 & 4 (0100) != 0 return true
func isCapture(m Move) bool{
	flag := GetFlag(m)
	return (flag & 4) != 0
}
// check for 4th bit 
// if 4th bit is 1 & 8 (1000) != 0 return true
func isPromotion(m Move) bool{
	flag := GetFlag(m)
	return (flag & 8) != 0
}