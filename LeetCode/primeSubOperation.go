package main

var minimumPrimeNumbers []int

func primeSubOperation(nums []int) bool {
	getPrimeNumbers(1000)

	for i := len(nums) - 2; i >= 0; i-- {
		diff := nums[i] - nums[i+1]
		if diff >= 0 {
			nums[i] -= minimumPrimeNumbers[diff]
			if nums[i] < 1 || nums[i] == nums[i+1] {
				return false
			}
		}
	}

	return true
}

func getPrimeNumbers(n int) {
	minimumPrimeNumbers = make([]int, n)
	pointer := 0
	for i := 2; i <= n; i++ {
		prime := true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			for k := pointer; k < i; k++ {
				minimumPrimeNumbers[k] = i
			}
			pointer = i
		}
	}
	for k := pointer; k < n; k++ {
		minimumPrimeNumbers[k] = pointer
	}
}
