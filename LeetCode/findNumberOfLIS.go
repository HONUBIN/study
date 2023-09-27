package main

import "math"

func findNumberOfLIS(nums []int) int {
	var n, ans, maxLength int = len(nums), 0, math.MinInt
	length := make([]int, n)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		length[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if length[i] < length[j]+1 {
					length[i] = length[j] + 1
					count[i] = 0
				}
				if length[i] == length[j]+1 {
					count[i] += count[j]
				}
			}
		}
		if length[i] > maxLength {
			maxLength = length[i]
			ans = count[i]
		} else if length[i] == maxLength {
			ans += count[i]
		}
	}
	return ans
}
