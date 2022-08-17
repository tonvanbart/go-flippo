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
		t.Error("a fresh created board should not be winning")
	}

	board[1][1] = true
	if !isWinning(board) {
		t.Error("a board with center set to true and all others to false is winning")
	}
}

func TestCornerMoves(t *testing.T) {
	board := [3][3]bool{{}}
	flip(&board, 1)
	if !board[0][0] || !board[0][1] || !board[1][0] || !board[1][1] {
		t.Error("Move '1' should flip the top left square")
	}

	board = [3][3]bool{{}}
	flip(&board, 3)
	if !board[0][1] || !board[0][2] || !board[1][1] || !board[1][2] {
		t.Error("Move '3' should flip the top right square")
	}

	board = [3][3]bool{{}}
	flip(&board, 7)
	if !board[1][0] || !board[1][1] || !board[2][0] || !board[2][1] {
		t.Error("Move '7' should flip the bottom left square")
	}

	board = [3][3]bool{{}}
	flip(&board, 9)
	if !board[1][1] || !board[1][2] || !board[2][1] || !board[2][2] {
		t.Error("Move '9' should flip the bottom right square")
	}
}

func TestSideMiddles(t *testing.T) {
	board := [3][3]bool{{}}
	flip(&board, 2)
	if !board[0][0] || !board[0][1] || !board[0][2] {
		t.Error("Move '2' should flip the top row")
	}

	board = [3][3]bool{{}}
	flip(&board, 8)
	if !board[2][0] || !board[2][1] || !board[2][2] {
		t.Error("Move '8' should flip the bottow row")
	}

	board = [3][3]bool{{}}
	flip(&board, 4)
	if !board[0][0] || !board[1][0] || !board[2][0] {
		t.Error("Move '4' should flip the left column")
	}

	board = [3][3]bool{{}}
	flip(&board, 6)
	if !board[0][2] || !board[1][2] || !board[2][2] {
		t.Error("Move '6' should flip the right column")
	}
}

func TestCenter(t *testing.T) {
	board := [3][3]bool{{}}
	flip(&board, 5)
	if (!board[0][1]) || !board[1][0] || !board[1][1] || !board[1][2] || !board[2][1] {
		t.Error("Move '5' should flip the center and the middles of all sides")
	}
}

func TestCompleteGame(t *testing.T) {
	board := [3][3]bool{{true, true, true}, {false, true, false}, {true, true, true}}
	printBoard(board)
	flip(&board, 2)
	printBoard(board)
	flip(&board, 8)
	printBoard(board)
	if !isWinning(board) {
		t.Error("Should have produced a win board")
	}
}
