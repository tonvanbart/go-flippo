## The ancient game of Flippo

This is my first Go program, a re-creation of a game I used to play many moons ago on a Unisys monochrome terminal.

From the instructions in the program:

The game is played on a 3x3 grid of cells, marked "." or "O". You start out with a random board, and your object is to get to a board where only the center cell is marked "O".
The cells on the board are numbered 1-9 and you play by typing the number of a cell to change.
As an example, a game in progress would look like this:
```

O . .   1 2 3
. O O   4 5 6
. O .   7 8 9

Choose a move 1-9, or 0 or <Enter> to quit:
```

Note that "flipping" a cell will also flip some of it's neighbours, following these rules:
 1. flipping a corner cell will flip all it's neigbours
 2. flipping the middle of a side will flip the entire side
 3. flipping the center will flip the middles of all sides.
