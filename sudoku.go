package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
)

// Sudoku represents a puzzle with a solution and identifying metadata
type Sudoku struct {
	ID       string  `json:"id"`
	Puzzle   [81]int `json:"puzzle"`
	Solution [81]int `json:"solution"`
	Name     string  `json:"name"`
}

func newSudoku(puzzle, solution [81]int) Sudoku {
	id := generateID(puzzle, solution)
	return Sudoku{
		ID:       id,
		Puzzle:   puzzle,
		Solution: solution,
		Name:     id,
	}
}

func generateID(puzzle, solution [81]int) string {
	var source bytes.Buffer

	for _, number := range puzzle {
		source.WriteByte(byte(number))
	}
	for _, number := range solution {
		source.WriteByte(byte(number))
	}

	hasher := sha1.New()
	hasher.Write(source.Bytes())
	return hex.EncodeToString(hasher.Sum(nil))
}
