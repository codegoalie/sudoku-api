package main

import (
	"encoding/json"
	"net/http"
)

func PuzzleShow(repo repo, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(repo.RandomSudoku()); err != nil {
		panic(err)
	}
}
