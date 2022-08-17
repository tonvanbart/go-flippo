package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// generate a new random board, making sure it's not a win
	board := [3][3]bool{{}}
	for board = randomBoard(); isWinning(board); {
		board = randomBoard()
	}
	printBoard(board)
	// testBoard := [3][3]bool{{true, false, false}, {false, true, false}, {true, true, true}}
	// printBoard(testBoard)
}

// randomBoard creates a new board with a random pattern.
func randomBoard() [3][3]bool {
	board := [3][3]bool{{}}
	rand.Seed(time.Now().UnixNano())
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if rand.Intn(2) == 1 {
				board[row][col] = true
			}
		}
	}
	return board
}

// check if a board is "winning", that is, if only the value at [1][1] is true and all the others are false.
func isWinning(board [3][3]bool) bool {
	return !board[0][0] && !board[0][1] && !board[0][2] &&
		!board[1][0] && board[1][1] && !board[1][2] &&
		!board[2][0] && !board[2][1] && !board[2][2]

}

// modify the board by flipping the indicated field and it's neighbours.
func flip(board [3][3]bool, move int) {
	if move == 1 {
		// stuff
	}

}

// flip a single field on the board, converting the move 1-9 to row and column coordinates.
// Here, the board is passed as a pointer and we modify the array in place.
func flipField(board *[3][3]bool, move int) {
	row := (move - 1) / 3
	col := (move - 1) % 3
	fmt.Printf("going to flip row=%d col=%d, before is %t\n", row, col, board[row][col])
	board[row][col] = !board[row][col]
	fmt.Printf("after flipping row=%d col=%d, value is now %t\n", row, col, board[row][col])
	// return board
}

// printBoard prints the board as a square of 'O' and '.'.
// After each row it prints the number of the corresponding move.
func printBoard(board [3][3]bool) {
	fmt.Println()
	for i, row := range board {
		for _, rowcol := range row {
			if rowcol {
				fmt.Print("O ")
			} else {
				fmt.Print(". ")
			}
		}
		offset := 3 * i
		fmt.Printf("   %d %d %d\n", offset+1, offset+2, offset+3)
	}
	fmt.Println()
}

// struct for (maybe) future use.
type flippo struct {
	fields [3][3]bool
}

// Stringer function to format the flippo struct for printing.
func (board flippo) String() string {
	return "flippo stringer (TO DO)"

}
