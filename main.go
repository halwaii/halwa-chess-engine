package main
import "fmt"
func main(){
	fmt.Println("welcome to halwa chess engine :P\n")
	board := [8][8]string{
		{"r","n","b","q","k","b","n","r"},
		{"p","p","p","p","p","p","p","p"},
		{".",".",".",".",".",".",".","."},
		{".",".",".",".",".",".",".","."},
		{".",".",".",".",".",".",".","."},
		{".",".",".",".",".",".",".","."},
		{"P","P","P","P","P","P","P","P"},
		{"R","N","B","Q","K","B","N","R"},
	}
	for i:=0;i<8;i++{
		for j:=0;j<8;j++{
			fmt.Printf("%v",board[i][j])
		}
		fmt.Println();
	}
}
