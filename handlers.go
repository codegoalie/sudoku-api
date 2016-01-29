package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type CreateParams struct {
	Puzzle   [81]int `json:"puzzle"`
	Solution [81]int `json:"solution"`
}

type ErrorBody struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type CreateErrorBody struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Id      string `json:"id"`
}

type Stats struct {
	Count int `json:"count"`
}

func StatsIndex(repo repo, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	count := repo.GetPuzzleCount()
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Stats{Count: count}); err != nil {
		panic(err)
	}
}

func PuzzleIndex(repo repo, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	sudoku, err := repo.RandomSudoku()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		body := ErrorBody{
			Error:   "Error fetching random puzzle",
			Message: err.Error(),
		}
		if err := json.NewEncoder(w).Encode(body); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(sudoku); err != nil {
			panic(err)
		}
	}
}

func PuzzleShow(repo repo, w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	id := parts[len(parts)-1]
	fmt.Println(id)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	sudoku, err := repo.GetSudoku(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		body := ErrorBody{
			Error:   "Error fetching puzzle",
			Message: err.Error(),
		}
		if err := json.NewEncoder(w).Encode(body); err != nil {
			panic(err)
		}
	} else {
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(sudoku); err != nil {
			panic(err)
		}
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

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	sudoku, err := repo.CreateSudoku(params.Puzzle, params.Solution)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		body := CreateErrorBody{
			Error:   "Error creating puzzle",
			Message: err.Error(),
			Id:      sudoku.Id,
		}
		if err.Error() == "Puzzle already exists" {
			duplicate.Inc()
		}
		if err := json.NewEncoder(w).Encode(body); err != nil {
			panic(err)
		}
	} else {
		added.Inc()
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(sudoku); err != nil {
			panic(err)
		}
	}
}
