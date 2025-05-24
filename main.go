package main

import (
	"fmt"
	"strings"
)

// TicTacToe represents the game board.
type TicTacToe struct {
	board [3][3]string
}

// NewTicTacToe initializes a new Tic Tac Toe game.
func NewTicTacToe() *TicTacToe {
	return &TicTacToe{board: [3][3]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}}
}

// Display prints the current state of the board.
func (t *TicTacToe) Display() {
	for _, row := range t.board {
		fmt.Println(strings.Join(row[:], "|"))
		if row != t.board[2] {
			fmt.Println("-----")
		}
	}
}

// MakeMove places a mark (X or O) at the specified position.
func (t *TicTacToe) MakeMove(row, col int, mark string) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 || t.board[row][col] != " " {
		return false
	}
	t.board[row][col] = mark
	return true
}

// CheckWinner determines if there is a winner.
func (t *TicTacToe) CheckWinner() string {
	for i := 0; i < 3; i++ {
		// Check rows and columns
		if t.board[i][0] != " " && t.board[i][0] == t.board[i][1] && t.board[i][1] == t.board[i][2] {
			return t.board[i][0]
		}
		if t.board[0][i] != " " && t.board[0][i] == t.board[1][i] && t.board[1][i] == t.board[2][i] {
			return t.board[0][i]
		}
	}
	// Check diagonals
	if t.board[0][0] != " " && t.board[0][0] == t.board[1][1] && t.board[1][1] == t.board[2][2] {
		return t.board[0][0]
	}
	if t.board[0][2] != " " && t.board[0][2] == t.board[1][1] && t.board[1][1] == t.board[2][0] {
		return t.board[0][2]
	}
	return " "
}

// IsBoardFull checks if the board is full.
func (t *TicTacToe) IsBoardFull() bool {
	for _, row := range t.board {
		for _, cell := range row {
			if cell == " " {
				return false
			}
		}
	}
	return true
}

func main() {
	game := NewTicTacToe()
	player := "X"

	for {
		game.Display()
		var row, col int
		fmt.Printf("Player %s, enter your move (row and column): ", player)
		fmt.Scan(&row, &col)

		if !game.MakeMove(row, col, player) {
			fmt.Println("Invalid move. Try again.")
			continue
		}

		winner := game.CheckWinner()
		if winner != " " {
			game.Display()
			fmt.Printf("Player %s wins!\n", winner)
			break
		}

		if game.IsBoardFull() {
			game.Display()
			fmt.Println("It's a draw!")
			break
		}

		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}
}
