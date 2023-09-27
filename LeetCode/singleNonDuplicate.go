package main

func singleNonDuplicate(nums []int) int {
	start := 0
	end := len(nums) - 1
	for {
		pivot := (start + end) / 2
		if start >= end {
			return nums[pivot]
		}
		if pivot%2 == 0 {
			if nums[pivot] == nums[pivot-1] {
				end = pivot - 2
			} else if nums[pivot] == nums[pivot+1] {
				start = pivot + 2
			} else {
				return nums[pivot]
			}
		} else {
			if nums[pivot] != nums[pivot-1] {
				end = pivot - 1
			} else if nums[pivot] != nums[pivot+1] {
				start = pivot + 1
			} else {
				return nums[pivot]
			}
		}
	}
}
