package leetcode

func climbStairs(n int) int {
	s1, s2 := 1, 2
	if n <= 2 {
		return n
	}

	for i := 3; i <= n; i++ {
		s := s1 + s2
		s1, s2 = s2, s
	}

	return s2
}
