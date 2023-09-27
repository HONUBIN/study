package main

import (
	"math"
	"strings"
)

func maxNumberOfBalloons(text string) int {
	balloon := map[rune]float64{
		'b': 0,
		'a': 0,
		'l': 0,
		'o': 0,
		'n': 0,
	}

	for _, c := range text {
		if strings.Contains("balloon", string(c)) {
			var cnt float64 = 1
			if c == 'l' || c == 'o' {
				cnt = 0.5
			}
			balloon[c] += cnt
		}
	}

	min := int(math.Floor(balloon['b']))
	for _, count := range balloon {
		if int(count) < min {
			min = int(count)
		}
	}

	return min
}
