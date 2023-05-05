package main

import (
	"fmt"
)

type Player struct {
	Name  string
	Score int
}

type Board struct {
	size [3][3]string
}

func (p Player) GetName() string {
	return p.Name
}

func (p Player) GetScore() int {
	return p.Score
}

func main() {
	player1 := new(Player)
	player2 := new(Player)

	fmt.Println("Enter player 1 name")
	fmt.Scanln(&player1.Name)
	fmt.Println("Enter player 2 name")
	fmt.Scanln(&player2.Name)

	var board Board
	var currentPlayer Player = *player1
	var winner string

	if currentPlayer.GetName() == player1.GetName() {
		currentPlayer = *player2
	} else {
		currentPlayer = *player1
	}

	gameMap := make(map[string]Player)
	gameMap[player1.GetName()] = *player1
	gameMap[player2.GetName()] = *player2

	for {
		fmt.Println(currentPlayer.GetName(), "turn")
		board.Print()

		x, y := GetMove()

		if board.IsValidMove(x, y) {
			board = board.MakeMove(x, y, currentPlayer.GetName())
			winner = board.CheckWinner()
			if winner != "" {
				break
			}
			if board.IsFull() {
				break
			}
		} else {
			fmt.Println("Invalid move")
		}
	}
}

func (b Board) Print() {
	fmt.Println("  0 1 2")
	for i := 0; i < len(b.size); i++ {
		fmt.Printf("%d %s\n", i, b.size[i])
	}
}

func (b Board) IsFull() bool {
	for i := 0; i < len(b.size); i++ {
		for j := 0; j < len(b.size); j++ {
			if b.size[i][j] == "" {
				return false
			}
		}
	}
	return true
}

func (b Board) IsValidMove(x, y int) bool {
	return x >= 0 && x < len(b.size) && y >= 0 && y < len(b.size) && b.size[x][y] == ""
}

func (b Board) CheckWinner() string {
	var winner string
	for i := 0; i < len(b.size); i++ {
		if b.size[i][0] == b.size[i][1] && b.size[i][1] == b.size[i][2] && b.size[i][0] != "" {
			winner = b.size[i][0]
			return winner
		}
		if b.size[0][i] == b.size[1][i] && b.size[1][i] == b.size[2][i] && b.size[0][i] != "" {
			winner = b.size[0][i]
			return winner
		}
	}
	// checking winner by diagonal and printing the player name
	if b.size[0][0] == b.size[1][1] && b.size[1][1] == b.size[2][2] && b.size[0][0] != "" {
		winner = b.size[0][0]
		return winner
	}
	if b.size[0][2] == b.size[1][1] && b.size[1][1] == b.size[2][0] && b.size[0][2] != "" {
		winner = b.size[0][2]
		return winner
	}

	return winner
}

func (b Board) MakeMove(x, y int, player string) Board {
	b.size[x][y] = player
	return b
}

func GetMove() (int, int) {
	var x, y int
	fmt.Println("Enter x coordinate")
	fmt.Scanln(&x)
	fmt.Println("Enter y coordinate")
	fmt.Scanln(&y)
	return x, y
}

func PlayAgain() bool {
	var playAgain string
	switch playAgain {
	case "y":
		return true
	case "n":
		return false
	default:
		fmt.Println("Invalid input")
	}
	fmt.Scan(&playAgain)

	if playAgain != "y" && playAgain != "n" {
		fmt.Println("Invalid input")
		return PlayAgain()
	}

	return playAgain == "y"
}
