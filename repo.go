package main

import (
	"fmt"

	"gopkg.in/redis.v3"
)

type repo interface {
	RandomSudoku() Sudoku
}

type RedisRepo struct {
	client *redis.Client
}

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
	return Sudoku{Id: "test"}
}
