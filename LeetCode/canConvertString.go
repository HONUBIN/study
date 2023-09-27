package main

func canConvertString(s string, t string, k int) bool {
	number := make([]int, 27)
	if len(s) != len(t) {
		return false
	}

	can := true
LOOP:
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			continue
		}
		move := int(t[i]) - int(s[i])
		if move < 0 {
			move += 26
		}
		if move+number[move] > k {
			can = false
			break LOOP
		}
		number[move] += 26
	}

	return can
}
