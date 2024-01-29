package main

import (
	"strconv"
)

func reverse(x int) int {
	s := strconv.Itoa(x)
	rs := ""
	if _, b := validate(s); !b {
		return 0
	} else if s[0] == '-' {
		rs = s[:1] + reverseString(s[1:])
	} else {
		rs = reverseString(s)
	}

	rx, b := validate(rs)
	if !b {
		return 0
	}

    return rx
}

func validate(s string) (int, bool) {
	x, err := strconv.ParseInt(s, 10, 32)
	return int(x), err == nil
}

func reverseString(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}