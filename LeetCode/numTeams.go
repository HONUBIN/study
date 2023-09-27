package main

func numTeams(rating []int) int {
	num := 0
	asc := make(map[int][]int, 0)
	desc := make(map[int][]int, 0)

	for i := 0; i < len(rating); i++ {
		for j := 0; j < i; j++ {
			if rating[i] > rating[j] {
				asc[rating[j]] = append(asc[rating[j]], rating[i])
			}
		}
	}

	for i := 0; i < len(rating); i++ {
		for j := 0; j < i; j++ {
			if rating[i] < rating[j] {
				desc[rating[j]] = append(desc[rating[j]], rating[i])
			}
		}
	}

	for _, i := range rating {
		for _, j := range asc[i] {
			num += len(asc[j])
		}

		for _, j := range desc[i] {
			num += len(desc[j])
		}
	}

	return num
}
