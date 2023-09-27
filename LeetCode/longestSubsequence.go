package main

func longestSubsequence(arr []int, difference int) int {
	count := make(map[int]int)
	for _, n := range arr {
		if val, ok := count[n-difference]; ok {
			count[n] = val + 1
		} else {
			count[n] = 1
		}
	}

	max := 1
	for _, c := range count {
		if c > max {
			max = c
		}
	}

	return max
}
