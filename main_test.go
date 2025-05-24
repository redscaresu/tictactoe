package main

import (
	"testing"
)

func TestMakeMove(t *testing.T) {
	game := NewGame()
	if !game.MakeMove(0, 0) {
		t.Error("Failed to make move at (0,0)")
	}
	if game.MakeMove(0, 0) {
		t.Error("Incorrectly allowed move at (0,0)")
	}
}

func TestCheckWin(t *testing.T) {
	game := NewGame()
	game.MakeMove(0, 0)
	game.MakeMove(0, 1)
	game.MakeMove(1, 1)
	game.MakeMove(0, 2)
	if game.CheckWin() != "X" {
		t.Error("Failed to detect win for X")
	}
}

func TestCheckDraw(t *testing.T) {
	game := NewGame()
	game.board = [3][3]string{
		{"X", "O", "X"},
		{"X", "X", "O"},
		{"O", "X", "O"},
	}
	if !game.CheckDraw() {
		t.Error("Failed to detect draw")
	}
}

func TestPlayGame(t *testing.T) {
	// This test is more of an integration test and would typically be run manually.
	game := NewGame()
	PlayGame(game)
}
