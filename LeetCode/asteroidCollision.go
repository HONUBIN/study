package main

// func asteroidCollision(asteroids []int) []int {
// 	for i := 1; i < len(asteroids); i++ {
// 		for j := i - 1; j >= 0; j-- {
// 			if asteroids[i] < 0 && asteroids[j] >= 0 {
// 				if -asteroids[i] > asteroids[j] {
// 					asteroids = append(asteroids[:j], asteroids[j+1:]...)
// 					i--
// 				} else if -asteroids[i] < asteroids[j] {
// 					asteroids = append(asteroids[:i], asteroids[i+1:]...)
// 					i--
// 					break
// 				} else {
// 					asteroids = append(asteroids[:i], asteroids[i+1:]...)
// 					i--
// 					asteroids = append(asteroids[:j], asteroids[j+1:]...)
// 					i--
// 					break
// 				}
// 			}
// 		}
// 	}

// 	return asteroids
// }

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	top := -1
	same := false
	for i := 0; i < len(asteroids); i++ {
		same = false
		asteroid := asteroids[i]

		if asteroids[i] > 0 {
			stack = append(stack, asteroid)
			top++
			continue
		}

		for {
			if same {
				break
			} else if len(stack) == 0 || asteroid > 0 || asteroid*stack[top] > 0 {
				stack = append(stack, asteroid)
				top++
				break
			} else {
				if -asteroid > stack[top] {
					stack = append(stack[:top], stack[top+1:]...)
				} else if -asteroid < stack[top] {
					asteroid = stack[top]
					stack = append(stack[:top], stack[top+1:]...)
				} else {
					stack = append(stack[:top], stack[top+1:]...)
					same = true
				}
				top--
			}
		}
	}

	return stack
}
