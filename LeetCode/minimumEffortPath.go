package main

import (
	"encoding/json"
	"math"
	"strconv"
)

var mem map[string]map[int]int
var globalHeights [][]int

func minimumEffortPath(heights [][]int) int {
	mem = make(map[string]map[int]int)
	globalHeights = heights
	path := make([][]int, len(heights))
	for i := range path {
		path[i] = make([]int, len(heights[0]))
	}

	return goAhead(0, 0, path, 0)
}

func goAhead(row int, col int, path [][]int, prevDirection int) int {
	min := 1000000
	pathBytes, _ := json.Marshal(path)
	pathString := strconv.Itoa(row) + strconv.Itoa(col) + string(pathBytes)
	if direction, ok := mem[pathString]; ok {
		for i := 0; i < 4; i++ {
			if i == prevDirection {
				continue
			}
			if direction[i] < min {
				min = direction[i]
			}
		}
	} else {
		for i := 0; i < 4; i++ {
			mem[pathString][i] = min
		}
	}

	if row == len(path)-1 && col == len(path[0])-1 {
		path[row][col] = 1
		return 0
	}

	path[row][col] = 1
	current := globalHeights[row][col]

	// up
	if row > 0 && col > 0 && col < len(path[0])-1 {
		i := row - 1
		j := col
		next := globalHeights[i][j]
		if path[i][j] != 1 {
			result := math.Max(math.Abs(float64(next-current)), float64(goAhead(i, j, path, 1)))
			mem[pathString][0] = int(result)
			path[i][j] = 0
			if int(result) < min {
				min = int(result)
			}
		}
	}

	// down
	if row < len(path)-1 {
		i := row + 1
		j := col
		next := globalHeights[i][j]
		if path[i][j] != 1 {
			result := math.Max(math.Abs(float64(next-current)), float64(goAhead(i, j, path, 0)))
			mem[pathString][1] = int(result)
			path[i][j] = 0
			if int(result) < min {
				min = int(result)
			}
		}
	}

	// left
	if row > 0 && row < len(path)-1 && col > 0 {
		i := row
		j := col - 1
		next := globalHeights[i][j]
		if path[i][j] != 1 {
			result := math.Max(math.Abs(float64(next-current)), float64(goAhead(i, j, path, 3)))
			mem[pathString][2] = int(result)
			path[i][j] = 0
			if int(result) < min {
				min = int(result)
			}
		}
	}

	// right
	if col < len(path[0])-1 {
		i := row
		j := col + 1
		next := globalHeights[i][j]
		if path[i][j] != 1 {
			result := math.Max(math.Abs(float64(next-current)), float64(goAhead(i, j, path, 2)))
			mem[pathString][3] = int(result)
			path[i][j] = 0
			if int(result) < min {
				min = int(result)
			}
		}
	}

	return min
}
