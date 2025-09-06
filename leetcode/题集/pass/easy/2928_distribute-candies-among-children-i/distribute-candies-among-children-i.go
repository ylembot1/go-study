package leetcode

func distributeCandies(n int, limit int) int {
	res := 0
	for i := 0; i <= limit; i++ {
		for j := 0; j <= limit; j++ {
			if i+j > n {
				break
			}
			if n-i-j <= limit {
				res++
			}
		}
	}

	return res
}
