package main

import (
	"flag"
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/codegoalie/sudoku-api/sudokuapi"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type server struct{}

type repo interface {
	GetSudoku(string) ([]uint32, error)
	GetCount() uint64
	CreatePuzzle(uuid string, grid []uint32) error
}

func (s server) GetPuzzle(ctx context.Context, params *pb.PuzzleID) (*pb.Puzzle, error) {
	grid, err := repoInstance.GetSudoku(params.Uuid)
	if err != nil {
		return nil, err
	}

	return &pb.Puzzle{Uuid: params.Uuid, Cell: grid}, nil
}

func (s server) GetStats(ctx context.Context, params *pb.StatsQuery) (*pb.Stats, error) {
	return &pb.Stats{Count: repoInstance.GetCount()}, nil
}

func (s server) CreatePuzzle(ctx context.Context, params *pb.Puzzle) (*pb.Puzzle, error) {
	err := repoInstance.CreatePuzzle(params.Uuid, params.Cell)
	return params, err
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", *port)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPuzzlerServer(grpcServer, server{})
	grpcServer.Serve(lis)
}
