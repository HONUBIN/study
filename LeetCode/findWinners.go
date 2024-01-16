package main

import (
	"sort"
)

func findWinners(matches [][]int) [][]int {
    winMap := make(map[int]int)
    looseMap := make(map[int]int)
	for _, match := range matches {
		winMap[match[0]]++
		looseMap[match[1]]++
	}
	// fmt.Println(winMap)
	// fmt.Println(looseMap)

	answer := make([][]int, 2)
	for key := range winMap {
		if _, ok := looseMap[key]; !ok {
			answer[0] = append(answer[0], key)
		}
	}
	for key, value := range looseMap {
		if value == 1 {
			answer[1] = append(answer[1], key)
		}
	}
	sort.Ints(answer[0])
	sort.Ints(answer[1])

	return answer
}
