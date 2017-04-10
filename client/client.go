package main

import (
	"context"
	"log"

	pb "github.com/codegoalie/sudoku-api/sudokuapi"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

const (
	address = "localhost:10000"
	id      = "3430eb79-7709-4957-bc63-0ceb650c4e6e"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewPuzzlerClient(conn)
	resp, err := client.GetPuzzle(context.Background(), &pb.PuzzleID{Uuid: id})
	if err != nil {
		log.Fatalf("Failed to get puzzle: %v", err)
	}

	log.Printf("Here's the pzzle: %v\n", resp)

	count, err := client.GetStats(context.Background(), &pb.StatsQuery{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Here's the count: %d\n", count.Count)

	puzzle := pb.Puzzle{Uuid: uuid.NewV4().String(), Start: []int32{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}}

	newPuzzle, err := client.CreatePuzzle(context.Background(), &puzzle)
	if err != nil {
		log.Fatalf("Count not create puzzle: %v", err)
	}
	log.Printf("Added: %v\n", newPuzzle.Uuid)

	log.Printf("Should fail:\n")
	newPuzzle, err = client.CreatePuzzle(context.Background(), &puzzle)
	if err != nil {
		log.Printf("Count not create puzzle: %v", err)
	}

}
