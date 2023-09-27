package main

func minSubArrayLen(target int, nums []int) int {
	startPoint := 0
	len := len(nums)
	minLen := len + 1
	sum := 0

	for i := 0; i < len; i++ {
		sum += nums[i]
		for {
			if sum-nums[startPoint] >= target {
				sum -= nums[startPoint]
				startPoint++
			} else {
				break
			}
		}
		if sum >= target {
			subLen := i - startPoint + 1
			if subLen < minLen {
				minLen = subLen
			}
		}
	}

	if minLen > len {
		return 0
	}
	return minLen
}
