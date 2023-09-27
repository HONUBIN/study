package main

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	sortIntervals(intervals)

	nonOverlappingIntervals := numberOfNonOverlappingIntervals(intervals)

	return len(intervals) - nonOverlappingIntervals
}

func numberOfNonOverlappingIntervals(intervals [][]int) int {
	currentInterval := intervals[0]
	var nbNonOverlappingIntervals int = 1

	for _, interval := range intervals[1:] {
		if interval[0] >= currentInterval[1] {
			currentInterval = interval
			nbNonOverlappingIntervals++
		}
	}

	return nbNonOverlappingIntervals
}

func sortIntervals(intervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
}
