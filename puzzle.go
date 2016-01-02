package main

type Sudoku struct {
	Id       string  `json:"id"`
	Puzzle   [81]int `json:"puzzle"`
	Solution [81]int `json:"solution"`
	Name     string  `json:"name"`
}
