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
	game.MakeMove(0, 0) // X at (0,0)
	game.MakeMove(1, 0) // O at (1,0)
	game.MakeMove(0, 1) // X at (0,1)
	game.MakeMove(1, 1) // O at (1,1)
	game.MakeMove(0, 2) // X at (0,2) - X wins top row
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

func TestBoundsChecking(t *testing.T) {
	game := NewGame()
	if game.MakeMove(-1, 0) {
		t.Error("Should not allow negative row")
	}
	if game.MakeMove(0, -1) {
		t.Error("Should not allow negative column")
	}
	if game.MakeMove(3, 0) {
		t.Error("Should not allow row >= 3")
	}
	if game.MakeMove(0, 3) {
		t.Error("Should not allow column >= 3")
	}
}
