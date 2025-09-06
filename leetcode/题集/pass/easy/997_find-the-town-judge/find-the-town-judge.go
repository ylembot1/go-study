package leetcode

type tr struct {
	trustOther int
	trusted    int
}

func findJudge(n int, trust [][]int) int {
	trL := make([]tr, n)
	for _, t := range trust {
		trL[t[0]-1].trustOther++
		trL[t[1]-1].trusted++
	}

	num := 0
	res := -1
	for idx, t := range trL {
		if t.trustOther == 0 && t.trusted == n-1 {
			num++
			res = idx + 1
		}
	}
	if num == 1 {
		return res
	}
	return -1
}
