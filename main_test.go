package main

import (
	"testing"
)

func TestMakeMove(t *testing.T) {
	game := NewTicTacToe()
	if !game.MakeMove(0, 0, "X") {
		t.Errorf("Expected move to be valid")
	}
	if game.MakeMove(0, 0, "O") {
		t.Errorf("Expected move to be invalid")
	}
	if game.MakeMove(3, 0, "X") {
		t.Errorf("Expected move to be invalid")
	}
}

func TestCheckWinner(t *testing.T) {
	game := NewTicTacToe()
	game.MakeMove(0, 0, "X")
	game.MakeMove(0, 1, "X")
	game.MakeMove(0, 2, "X")
	if game.CheckWinner() != "X" {
		t.Errorf("Expected winner to be X")
	}

	game = NewTicTacToe()
	game.MakeMove(0, 0, "O")
	game.MakeMove(1, 0, "O")
	game.MakeMove(2, 0, "O")
	if game.CheckWinner() != "O" {
		t.Errorf("Expected winner to be O")
	}

	game = NewTicTacToe()
	game.MakeMove(0, 0, "X")
	game.MakeMove(1, 1, "X")
	game.MakeMove(2, 2, "X")
	if game.CheckWinner() != "X" {
		t.Errorf("Expected winner to be X")
	}

	game = NewTicTacToe()
	game.MakeMove(0, 2, "O")
	game.MakeMove(1, 1, "O")
	game.MakeMove(2, 0, "O")
	if game.CheckWinner() != "O" {
		t.Errorf("Expected winner to be O")
	}
}

func TestIsBoardFull(t *testing.T) {
	game := NewTicTacToe()
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			game.MakeMove(row, col, "X")
		}
	}
	if !game.IsBoardFull() {
		t.Errorf("Expected board to be full")
	}

	game = NewTicTacToe()
	if game.IsBoardFull() {
		t.Errorf("Expected board to not be full")
	}
}
