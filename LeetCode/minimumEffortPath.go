package main

var efforts [][]int
var visit [][]bool
var rows int
var cols int
var max int = 1000001

// queue method will improve the performance
func getMinNode() (int, int) {
	min := max
	row := -1
	col := -1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visit[i][j] && efforts[i][j] < min {
				min = efforts[i][j]
				row = i
				col = j
			}
		}
	}

	return row, col
}

func minimumEffortPath(heights [][]int) int {
	rows = len(heights)
	cols = len(heights[0])
	if rows == 1 && cols == 1 {
		return 0
	}

	efforts = make([][]int, rows)
	for i := range efforts {
		efforts[i] = make([]int, cols)
		for j := range efforts[i] {
			efforts[i][j] = max
		}
	}
	efforts[0][0] = 0

	visit = make([][]bool, rows)
	for i := range visit {
		visit[i] = make([]bool, cols)
	}

	for {
		i, j := getMinNode()
		if i < 0 && j < 0 {
			break
		}
		visit[i][j] = true
		update(heights, i, j, rows, cols)
	}

	return efforts[rows-1][cols-1]
}

func update(heights [][]int, i, j, rows, cols int) {
	var diff int
	// left
	if j-1 > -1 {
		if abs(heights[i][j]-heights[i][j-1]) > efforts[i][j] {
			diff = abs(heights[i][j] - heights[i][j-1])
		} else {
			diff = efforts[i][j]
		}

		if diff < efforts[i][j-1] {
			efforts[i][j-1] = diff
		}
	}

	// right
	if j+1 < cols {
		if abs(heights[i][j]-heights[i][j+1]) > efforts[i][j] {
			diff = abs(heights[i][j] - heights[i][j+1])
		} else {
			diff = efforts[i][j]
		}

		if diff < efforts[i][j+1] {
			efforts[i][j+1] = diff
		}
	}

	// up
	if i-1 > -1 {
		if abs(heights[i][j]-heights[i-1][j]) > efforts[i][j] {
			diff = abs(heights[i][j] - heights[i-1][j])
		} else {
			diff = efforts[i][j]
		}

		if diff < efforts[i-1][j] {
			efforts[i-1][j] = diff
		}
	}

	// down
	if i+1 < rows {
		if abs(heights[i][j]-heights[i+1][j]) > efforts[i][j] {
			diff = abs(heights[i][j] - heights[i+1][j])
		} else {
			diff = efforts[i][j]
		}

		if diff < efforts[i+1][j] {
			efforts[i+1][j] = diff
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
