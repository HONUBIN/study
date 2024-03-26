package main

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	heaters = append([]int{math.MinInt64}, heaters...)
	heaters = append(heaters, math.MaxInt64)

	radius := 0
	for _, house := range houses {
		left := 0
		right := len(heaters) - 1

		for left <= right {
			mid := left + (right-left)/2
			if heaters[mid] < house {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		currRadius := abs(heaters[left] - house)
		if left-1 >= 0 {
			currRadius = findMin(currRadius, abs(house-heaters[left-1]))
		}
		radius = findMax(radius, currRadius)
	}

	return radius
}

// func abs(n int) int {
// 	if n < 0 {
// 		return -n
// 	}
// 	return n
// }

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
