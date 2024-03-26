package main

func findUnsortedSubarray(nums []int) int {
	end := 0
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < max {
			end = i
		} else {
			max = nums[i]
		}
	}

	start := len(nums)
	min := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > min {
			start = i
		} else {
			min = nums[i]
		}
	}

	if start > end {
		return 0
	}

	return end - start + 1
}
