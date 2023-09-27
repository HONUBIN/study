package main

import (
	"encoding/json"
)

var mem2 map[string]bool

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}

	if (maxChoosableInteger+1)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	mem2 = map[string]bool{}
	integers := make([]int, maxChoosableInteger+1)

	return takeTurns(maxChoosableInteger, desiredTotal, integers)
}

func takeTurns(maxChoosableInteger int, desiredTotal int, integers []int) bool {
	integersString, _ := json.Marshal(integers)
	if value, ok := mem2[string(integersString)]; ok {
		return value
	}

	for i := 1; i <= maxChoosableInteger; i++ {
		if integers[i] == 1 {
			continue
		}

		integers[i] = 1
		if desiredTotal <= i || !takeTurns(maxChoosableInteger, desiredTotal-i, integers) {
			mem2[string(integersString)] = true
			integers[i] = 0
			return true
		}
		integers[i] = 0
	}

	mem2[string(integersString)] = false
	return false
}
