package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/redis.v3"
)

type repo interface {
	RandomSudoku() Sudoku
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
		return Sudoku{Id: "error"}
	}
	if err = json.Unmarshal(bs, &sudoku); err != nil {
		fmt.Println(err)
		return Sudoku{Id: "error"}
	}
	return sudoku
}
