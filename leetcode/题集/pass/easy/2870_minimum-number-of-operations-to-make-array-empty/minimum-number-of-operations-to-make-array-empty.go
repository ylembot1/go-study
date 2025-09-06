package leetcode

func minOperations(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	dp := make([]int, 100001)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	dp[2] = 1
	dp[3] = 1
	for i := 4; i < len(dp); i++ {
		t2 := dp[i-2]
		t3 := dp[i-3]
		t := t2
		if t == -1 || t > t3 && t3 != -1 {
			t = t3
		}

		if t != -1 {
			dp[i] = t + 1
		}
	}

	res := 0
	for _, v := range m {
		if dp[v] == -1 {
			return -1
		}
		res += dp[v]
	}

	return res
}
