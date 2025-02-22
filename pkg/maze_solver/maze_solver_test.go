package mazesolver

import (
	"strings"
	"testing"

	"github.com/samber/lo"
)

func TestMazeSolver(t *testing.T) {
	maze := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	mazeResult := []Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
		{x: 1, y: 5},
	}
	result := MazeSolver(maze, "x", Point{x: 10, y: 0}, Point{x: 1, y: 5})

	if drawPath(maze, result) != drawPath(maze, mazeResult) {
		t.Errorf("Expected %v, got %v", result, mazeResult)
	}

}

func drawPath(data []string, path []Point) string {
	data2 := lo.Map(data, func(item string, _ int) []string {
		return strings.Split(item, "")
	})
	for _, p := range path {
		if len(data2[p.y]) != 0 && data2[p.y][p.x] != " " {
			data2[p.y][p.x] = "*"
		}
	}
	return strings.Join(lo.Map(data2, func(item []string, _ int) string {
		return strings.Join(item, "")
	}), "")
}
