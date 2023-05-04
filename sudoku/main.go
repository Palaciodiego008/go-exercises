package main

import "fmt"

const boardSize = 9

func main() {
	var board [boardSize][boardSize]int

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			fmt.Printf("Enter value for board[%d][%d]: ", i, j)
			fmt.Scanf("%d", &board[i][j])
		}
	}
	printBoard(board)
}

func printBoard(board [boardSize][boardSize]int) {
	for i := 0; i < len(board); i++ {
		if i%3 == 0 {
			fmt.Println(" -----------------------")
		}
		for j := 0; j < len(board[i]); j++ {
			if j%3 == 0 {
				fmt.Print("| ")
			}
			fmt.Print(board[i][j])
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
	fmt.Println(" -----------------------")

}
