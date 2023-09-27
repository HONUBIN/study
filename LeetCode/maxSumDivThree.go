package main

import (
	"sort"
)

func maxSumDivThree(nums []int) int {
	sum := 0
	remainder_1 := make([]int, 0)
	remainder_2 := make([]int, 0)
	for _, num := range nums {
		sum += num
		if num%3 == 1 {
			remainder_1 = append(remainder_1, num)
		} else if num%3 == 2 {
			remainder_2 = append(remainder_2, num)
		}
	}

	sort.Ints(remainder_1)
	sort.Ints(remainder_2)

	if sum%3 == 1 {
		if len(remainder_1) < 1 {
			sum -= remainder_2[0] + remainder_2[1]
		} else if len(remainder_2) < 2 || remainder_1[0] < remainder_2[0]+remainder_2[1] {
			sum -= remainder_1[0]
		} else {
			sum -= remainder_2[0] + remainder_2[1]
		}
	} else if sum%3 == 2 {
		if len(remainder_2) < 1 {
			sum -= remainder_1[0] + remainder_1[1]
		} else if len(remainder_1) < 2 || remainder_2[0] < remainder_1[0]+remainder_1[1] {
			sum -= remainder_2[0]
		} else {
			sum -= remainder_1[0] + remainder_1[1]
		}
	}

	return sum
}
