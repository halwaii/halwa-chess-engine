package main

// move encoding 
type Move uint16
// 2^6 = 64
// we need 6 bits to define each square
// 0 to 5 bits - from
// 6 to 11 bits - to
// 12 to 15 bits - flag
// func to encode move
func EncodeMove(from uint16, to uint16, flag uint16) Move{
	return Move(from | (to << 6) | (flag << 12))
}
// move decode
func (m Move) GetFrom() uint16{
	return uint16(m) & 0x3F // 0x3F is 63
}
func (m Move) GetTo() uint16{
	return (uint16(m) >> 6) & 0x3F 
}
func (m Move) GetFlag() uint16{
	return (uint16(m) >> 12) & 0x3F
}