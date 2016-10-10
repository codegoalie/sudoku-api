package main

import "errors"

var (
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

func (r staticRepo) GetSudoku(uuid string) ([]uint32, error) {
	if grid, ok := r.db[uuid]; ok {
		return grid, nil
	}

	return []uint32{}, errors.New("Unknown puzzle UUID")
}

func (r staticRepo) GetCount() uint64 {
	return uint64(len(r.db))
}

func (r staticRepo) CreatePuzzle(uuid string, grid []uint32) error {
	if len(r.db[uuid]) > 0 {
		return errors.New("Puzzle already exists")
	}
	r.db[uuid] = grid
	return nil
}
