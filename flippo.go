package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// main logic
func main() {
	// generate a new random board, making sure it's not a win
	board := [3][3]bool{{}}
	for board = randomBoard(); isWinning(board); {
		board = randomBoard()
	}

	var yesno string
	fmt.Print("Welcome to the ancient game of Flippo! Would you like instructions (y/n)? ")
	fmt.Scanf("%s", &yesno)
	if yesno == "y" || yesno == "Y" {
		printHelp()
	}
	printBoard(board)

	for !isWinning(board) {
		move := getMove()
		flip(&board, move)
		printBoard(board)
	}

	fmt.Println("You won! Congratulations!")
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
// Note that the array is changed in-place.
func flip(board *[3][3]bool, move int) {
	if move == 1 {
		flipField(board, 1)
		flipField(board, 2)
		flipField(board, 4)
		flipField(board, 5)
	} else if move == 3 {
		flipField(board, 2)
		flipField(board, 3)
		flipField(board, 5)
		flipField(board, 6)
	} else if move == 7 {
		flipField(board, 4)
		flipField(board, 5)
		flipField(board, 7)
		flipField(board, 8)
	} else if move == 9 {
		flipField(board, 5)
		flipField(board, 6)
		flipField(board, 8)
		flipField(board, 9)
	} else if move == 2 {
		flipField(board, 1)
		flipField(board, 2)
		flipField(board, 3)
	} else if move == 8 {
		flipField(board, 7)
		flipField(board, 8)
		flipField(board, 9)
	} else if move == 4 {
		flipField(board, 1)
		flipField(board, 4)
		flipField(board, 7)
	} else if move == 6 {
		flipField(board, 3)
		flipField(board, 6)
		flipField(board, 9)
	} else if move == 5 {
		flipField(board, 2)
		flipField(board, 4)
		flipField(board, 5)
		flipField(board, 6)
		flipField(board, 8)
	}
}

// flip a single field on the board, converting the move 1-9 to row and column coordinates.
// Here, the board is passed as a pointer and we modify the array in place.
func flipField(board *[3][3]bool, move int) {
	row := (move - 1) / 3
	col := (move - 1) % 3
	board[row][col] = !board[row][col]
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

// Prompt and gets the next move making sure that it's a number 1-9 or quit if it's 0.
func getMove() int {
	var move int
	for {
		fmt.Print("Choose a move 1-9, or 0 or <Enter> to quit: ")
		fmt.Scanf("%d", &move)
		if move == 0 {
			fmt.Println("Bye!")
			os.Exit(0)
		}
		if move > 0 && move < 10 {
			return move
		}
	}

}

// Print instructions for the program.
func printHelp() {
	instructions := `
This is the ancient terminal game of Flippo.
The game is played on a 3x3 grid of cells, marked "." or "O". You start out with a random
board, and your object is to get to a board where only the center cell is marked "O".
The cells on the board are numbered 1-9 and you play by typing the number of a cell to change.
An example board would look like this:

O . .   1 2 3
. O O   4 5 6
. O .   7 8 9

Note that "flipping" a cell will also flip some of it's neighbours, following these rules:
 1. flipping a corner cell will flip all it's neigbours
 2. flipping the middle of a side will flip the entire side
 3. flipping the center will flip the middles of all sides.

 Your game starts below, have fun!
	`
	fmt.Print(instructions)
}

// struct for (maybe) future use.
type flippo struct {
	fields [3][3]bool
}

// Stringer function to format the flippo struct for printing.
func (board flippo) String() string {
	return "flippo stringer (TO DO)"

}
