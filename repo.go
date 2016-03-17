package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/redis.v3"
)

type repo interface {
	GetSudoku(string) (Sudoku, error)
	RandomSudoku() (Sudoku, error)
	CreateSudoku([81]int, [81]int) (Sudoku, error)
	GetPuzzleCount() int
}

type redisRepo struct {
	client *redis.Client
}

var puzzlesKey = "puzzles"
var listKey = "puzzleList"

func newRedisRepo(addr string) redisRepo {
	fmt.Println("Connecting to redis on", addr)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return redisRepo{client: client}
}

func (r redisRepo) GetPuzzleCount() int {
	return int(r.client.SCard(listKey).Val())
}

func (r redisRepo) RandomSudoku() (Sudoku, error) {
	id, err := r.client.SRandMember(listKey).Bytes()
	if err != nil {
		return Sudoku{}, errors.New("No puzzles in list")
	}
	return r.GetSudoku(string(id))
}

func (r redisRepo) GetSudoku(id string) (Sudoku, error) {
	var sudoku Sudoku
	bs, err := r.client.Get(puzzlesKey + id).Bytes()
	if err != nil {
		return Sudoku{}, errors.New("No puzzle found error")
	}
	if err = json.Unmarshal(bs, &sudoku); err != nil {
		return Sudoku{}, errors.New("Puzzle unmarshal error")
	}
	return sudoku, nil
}

func (r redisRepo) CreateSudoku(puzzle, solution [81]int) (Sudoku, error) {
	sudoku := newSudoku(puzzle, solution)

	_, err := r.GetSudoku(sudoku.ID)
	if err == nil {
		return sudoku, errors.New("Puzzle already exists")
	}

	b, err := json.Marshal(sudoku)
	if err != nil {
		fmt.Println(err)
		return sudoku, errors.New("Cannot parse params")
	}

	r.client.SAdd(listKey, sudoku.ID)
	r.client.Set(puzzlesKey+sudoku.ID, string(b), 0)
	if err != nil {
		fmt.Println(err)
		return sudoku, errors.New("Failed to persist puzzle")
	}

	return sudoku, nil
}
