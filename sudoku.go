package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
)

type Sudoku struct {
	Id       string  `json:"id"`
	Puzzle   [81]int `json:"puzzle"`
	Solution [81]int `json:"solution"`
	Name     string  `json:"name"`
}

func NewSudoku(puzzle, solution [81]int) Sudoku {
	id := generateId(puzzle, solution)
	return Sudoku{
		Id:       id,
		Puzzle:   puzzle,
		Solution: solution,
		Name:     id,
	}
}

func generateId(puzzle, solution [81]int) string {
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
