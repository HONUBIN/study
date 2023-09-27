package main

func longestArithSeqLength(nums []int) int {
	arithMap := make(map[int]map[int]int, 0)
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if arithMap[i] == nil {
				arithMap[i] = map[int]int{nums[i] - nums[j]: arithMap[j][nums[i]-nums[j]] + 1}
			} else {
				if arithMap[i][nums[i]-nums[j]] < arithMap[j][nums[i]-nums[j]]+1 {
					arithMap[i][nums[i]-nums[j]] = arithMap[j][nums[i]-nums[j]] + 1
				}
			}
		}
	}

	max := 0
	for _, subMap := range arithMap {
		for _, v := range subMap {
			if v > max {
				max = v
			}
		}
	}

	return max + 1
}
