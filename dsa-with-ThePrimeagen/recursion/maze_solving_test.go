package recursion

import (
	"fmt"
	"testing"
)

func TestMazeSolving(t *testing.T) {
	maze := [][]string{
		{"x", "x", "x", " ", "x"},
		{"x", "x", "x", " ", "x"},
		{"x", "x", "x", " ", "x"},
		{"x", "x", "x", " ", "x"},
		{"x", " ", " ", " ", "x"},
		{"x", " ", "x", "x", "x"},
		{"x", " ", "x", "x", "x"},
	}

	res := []Point{
		{x: 3, y: 0},
		{x: 3, y: 1},
		{x: 3, y: 2},
		{x: 3, y: 3},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
		{x: 1, y: 6},
	}

	path := maze_solve(maze, "x", Point{x: 3, y: 0}, Point{x: 1, y: 6})
	fmt.Println(path)

	for i, v := range res {
		if v.x != path[i].x || v.y != path[i].y {
			t.Fatalf("Failed expected {x=%d,y=%d} but rec {x=%d,y=%d}", v.x, v.y, path[i].x, path[i].y)
		}
	}
}
