package main

type node struct {
	believeOthers int
	believed      int
}

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}

	allPeople := make([]node, n+1)

	for _, edge := range trust {
		allPeople[edge[0]].believeOthers += 1
		allPeople[edge[1]].believed += 1
	}

	res := -1
	for idx, people := range allPeople {
		if people.believeOthers == 0 && people.believed == n-1 {
			if res != -1 {
				return -1
			}
			res = idx
		}
	}
	return res
}
