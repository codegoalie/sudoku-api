# sudoku-api

[ ![Codeship Status for chrismar035/sudoku-api](https://codeship.com/projects/e65c0d10-a3ba-0133-077d-1ac8bff03ae9/status?branch=master)](https://codeship.com/projects/129316)

The HTTP api service for serving Sudoku puzzles

`GET /puzzle`

```
{
  "id": "d49f42a75566e5315766231b3dcd5ef93e7fd063",
  "puzzle": [
    0, 3, 0, 0, 0, 0, 0, 5, 0,
    0, 0, 8, 0, 9, 1, 3, 0, 0,
    6, 0, 0, 4, 0, 0, 7, 0, 0,
    0, 0, 3, 8, 1, 0, 0, 0, 0,
    0, 0, 6, 0, 0, 0, 2, 0, 0,
    0, 0, 0, 0, 3, 4, 8, 0, 0,
    0, 0, 1, 0, 0, 8, 0, 0, 9,
    0, 0, 4, 1, 2, 0, 6, 0, 0,
    0, 6, 0, 0, 0, 0, 0, 4, 0
  ],
  "solution": [
    1, 3, 9, 7, 6, 2, 4, 5, 8,
    7, 4, 8, 5, 9, 1, 3, 2, 6,
    6, 5, 2, 4, 8, 3, 7, 9, 1,
    5, 2, 3, 8, 1, 6, 9, 7, 4,
    4, 8, 6, 9, 5, 7, 2, 1, 3,
    9, 1, 7, 2, 3, 4, 8, 6, 5,
    2, 7, 1, 6, 4, 8, 5, 3, 9,
    3, 9, 4, 1, 2, 5, 6, 8, 7,
    8, 6, 5, 3, 7, 9, 1, 4, 2
  ],
  "name": "Angry Elf"
}
```
