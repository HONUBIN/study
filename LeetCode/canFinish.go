package main

var courseLink [][]int
var status []int

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	courseLink = make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		courseLink[prerequisite[0]] = append(courseLink[prerequisite[0]], prerequisite[1])
	}

	// -1 = visiting, 0 = pending, 1 = can, 2 = can't
	status = make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		dfs(i)
		if status[i] == 2 {
			return false
		}
	}

	return true
}

func dfs(i int) {
	if status[i] != 0 {
		return
	}

	status[i] = -1
	s := 1
	for _, j := range courseLink[i] {
		dfs(j)
		if status[j] != 1 {
			s = 2
			break
		}
	}
	status[i] = s
}
