package main

func majorityElement(nums []int) []int {
	var elements []int
	counts := make(map[int]int)
	standard := len(nums) / 3

	for _, num := range nums {
		counts[num]++
		if counts[num] > standard {
			elements = append(elements, num)
		}
	}

	return elements
}
