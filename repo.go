package main

type repo interface {
	getSudoku(string) (Sudoku, error)
	randomSudoku() (Sudoku, error)
	createSudoku([81]int, [81]int) (Sudoku, error)
	getPuzzleCount() int
}
