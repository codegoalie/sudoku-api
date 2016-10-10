package main

import (
	"errors"
	"flag"
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/codegoalie/sudoku-api/sudokuapi"
)

var (
	port         = flag.Int("port", 10000, "The server port")
	repoInstance = staticRepo{db: map[string][]uint32{
		"3430eb79-7709-4957-bc63-0ceb650c4e6e": []uint32{
			7, 5, 0, 0, 0, 0, 0, 2, 0,
			1, 0, 0, 2, 0, 0, 0, 0, 0,
			3, 0, 0, 0, 9, 0, 4, 0, 6,
			0, 0, 0, 1, 7, 0, 0, 0, 0,
			0, 0, 1, 0, 3, 0, 5, 0, 0,
			0, 0, 0, 0, 4, 8, 0, 0, 0,
			8, 0, 9, 0, 5, 0, 0, 0, 2,
			0, 0, 0, 0, 0, 7, 0, 0, 3,
			0, 6, 0, 0, 0, 0, 0, 5, 1,
		},
	},
	}
)

type staticRepo struct {
	db map[string][]uint32
}

type server struct{}

type repo interface {
	GetSudoku(string) ([]uint32, error)
	GetCount() uint64
}

func (r staticRepo) GetSudoku(uuid string) ([]uint32, error) {
	if grid, ok := r.db[uuid]; ok {
		return grid, nil
	}

	return []uint32{}, errors.New("Unknown puzzle UUID")
}

func (r staticRepo) GetCount() uint64 {
	return uint64(len(r.db))
}

func (s server) GetPuzzle(ctx context.Context, params *pb.PuzzleID) (*pb.Puzzle, error) {
	fmt.Println(params.Uuid)
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
	if len(repoInstance.db[params.Uuid]) > 0 {
		return nil, errors.New("Puzzle already exists")
	}
	repoInstance.db[params.Uuid] = params.Cell
	return params, nil
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
