package main

func countPalindromicSubsequence(s string) int {
	left := make(map[rune]int, 26)
	right := make(map[rune]int, 26)
	for _, char := range s {
		right[char]++
	}
	subsequences := make(map[[3]rune]bool)
	for _, char := range s {
		right[char]--
		for j := 'a'; j <= 'z'; j++ {
			if left[j] > 0 && right[j] > 0 {
				subsequences[[3]rune{j, char, j}] = true
			}
		}
		left[char]++
	}
	return len(subsequences)
}
