package main

import (
	"context"
	"log"

	pb "github.com/codegoalie/sudoku-api/sudokuapi"
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

	log.Printf("Here's the pzzle: %v", resp)

	count, err := client.GetStats(context.Background(), &pb.StatsQuery{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Here's the count: %d", count.Count)
}
