package main

import "testing"

func TestFlipField(t *testing.T) {
	testboard := [3][3]bool{{}}
	// testboard = flipField(testboard, 2)
	flipField(&testboard, 2)
	printBoard(testboard)
	if !testboard[0][1] {
		t.Log("field [0][1] should have been set to true")
		t.Fail()
	}

}

func TestIsWinning(t *testing.T) {
	board := [3][3]bool{{}}
	if isWinning(board) {
		t.Log("a fresh created board should not be winning")
		t.Fail()
	}

	board[1][1] = true
	if !isWinning(board) {
		t.Log("a board with center set to true and all others to false is winning")
		t.Fail()
	}
}
