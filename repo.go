package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/redis.v3"
)

type repo interface {
	RandomSudoku() Sudoku
	CreateSudoku([81]int, [81]int) Sudoku
}

type RedisRepo struct {
	client *redis.Client
}

var puzzlesKey string = "puzzles"

func NewRedisRepo(addr string) RedisRepo {
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

	return RedisRepo{client: client}
}

func (r RedisRepo) RandomSudoku() Sudoku {
	var sudoku Sudoku
	bs, err := r.client.SRandMember(puzzlesKey).Bytes()
	if err != nil {
		fmt.Println(err)
		return Sudoku{Id: "No puzzle found error"}
	}
	if err = json.Unmarshal(bs, &sudoku); err != nil {
		fmt.Println(err)
		return Sudoku{Id: "Puzzle unmarshal error"}
	}
	return sudoku
}

func (r RedisRepo) CreateSudoku(puzzle, solution [81]int) Sudoku {
	sudoku := NewSudoku(puzzle, solution)

	b, err := json.Marshal(sudoku)
	if err != nil {
		fmt.Println(err)
		return Sudoku{Name: "Cannot parse params"}
	}

	r.client.SAdd(puzzlesKey, string(b))
	if err != nil {
		fmt.Println(err)
		return Sudoku{Name: "Failed to persist puzzle"}
	}

	return sudoku
}
