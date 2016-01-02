package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type CreateParams struct {
	Puzzle   [81]int `json:"puzzle"`
	Solution [81]int `json:"solution"`
}

func PuzzleShow(repo repo, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(repo.RandomSudoku()); err != nil {
		panic(err)
	}
}

func PuzzleCreate(repo repo, w http.ResponseWriter, r *http.Request) {
	var params CreateParams
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &params); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	fmt.Println("params:", params)

	sudoku := repo.CreateSudoku(params.Puzzle, params.Solution)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(sudoku); err != nil {
		panic(err)
	}
}
