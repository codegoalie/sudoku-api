package main

import (
	"context"
	"log"

	"google.golang.org/cloud/datastore"
)

type dsRepo struct {
	client *datastore.Client
}

// Grid holds some cells
type Grid struct {
	cells []uint32
}

func newDataStoreRepo(ctx context.Context, projectID string) *dsRepo {
	dsClient, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to connect to datastore: %v\n", err)
	}
	return &dsRepo{client: dsClient}
}

func (r dsRepo) GetSudoku(ctx context.Context, uuid string) ([]uint32, error) {
	var grid Grid
	key := datastore.NewKey(ctx, "puzzle", uuid, 0, nil)
	if err := r.client.Get(ctx, key, &grid); err != nil {
		return []uint32{}, err
	}

	return grid.cells, nil
}

func (r dsRepo) GetCount(ctx context.Context) uint64 {
	return 1
}

func (r dsRepo) CreatePuzzle(ctx context.Context, uuid string, grid []uint32) error {
	// if len(r.db[uuid]) > 0 {
	// 	return errors.New("Puzzle already exists")
	// }
	// r.db[uuid] = grid
	return nil
}
