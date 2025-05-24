package main

import (
	"fmt"
)

// Game represents the Tic Tac Toe game board and state.
type Game struct {
	board [3][3]string
	turn  string
}

// NewGame initializes a new Tic Tac Toe game.
func NewGame() *Game {
	return &Game{
		board: [3][3]string{
			{" ", " ", " "},
			{" ", " ", " "},
			{" ", " ", " "},
		},
		turn: "X",
	}
}

// PrintBoard prints the current state of the game board.
func (g *Game) PrintBoard() {
	for _, row := range g.board {
		fmt.Println(row)
	}
}

// MakeMove makes a move on the game board.
func (g *Game) MakeMove(row, col int) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return false // Out of bounds
	}
	if g.board[row][col] != " " {
		return false // Invalid move
	}
	g.board[row][col] = g.turn
	if g.turn == "X" {
		g.turn = "O"
	} else {
		g.turn = "X"
	}
	return true
}

// CheckWin checks if there is a winner.
func (g *Game) CheckWin() string {
	// Check rows and columns
	for i := 0; i < 3; i++ {
		if g.board[i][0] != " " && g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2] {
			return g.board[i][0]
		}
		if g.board[0][i] != " " && g.board[0][i] == g.board[1][i] && g.board[1][i] == g.board[2][i] {
			return g.board[0][i]
		}
	}
	// Check diagonals
	if g.board[0][0] != " " && g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] {
		return g.board[0][0]
	}
	if g.board[0][2] != " " && g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] {
		return g.board[0][2]
	}
	return ""
}

// CheckDraw checks if the game is a draw.
func (g *Game) CheckDraw() bool {
	for _, row := range g.board {
		for _, cell := range row {
			if cell == " " {
				return false
			}
		}
	}
	return true
}

// PlayGame starts the game loop.
func PlayGame(g *Game) {
	for {
		g.PrintBoard()
		var row, col int
		fmt.Printf("Player %s, enter your move (row and column): ", g.turn)
		fmt.Scan(&row, &col)
		if !g.MakeMove(row, col) {
			fmt.Println("Invalid move. Try again.")
			continue
		}
		winner := g.CheckWin()
		if winner != "" {
			g.PrintBoard()
			fmt.Printf("Player %s wins!\n", winner)
			break
		}
		if g.CheckDraw() {
			g.PrintBoard()
			fmt.Println("The game is a draw!")
			break
		}
	}
}

func main() {
	game := NewGame()
	PlayGame(game)
}
