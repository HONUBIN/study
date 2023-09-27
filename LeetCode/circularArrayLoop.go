package main

func circularArrayLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] < -1000 {
			continue
		}
		index := -(1001 + i)
		prev := i
		curr := i
		direction := 1
		if nums[i] < 0 {
			direction = -1
		}
		for {
			value := nums[curr]
			if value < -1000 {
				if value != index {
					break
				} else {
					return true
				}
			}
			if direction*value < 0 {
				break
			}
			nums[curr] = index
			prev = curr
			curr = (curr + value) % len(nums)
			if curr < 0 {
				curr += len(nums)
			}
			if prev == curr {
				break
			}
		}
	}

	return false
}
