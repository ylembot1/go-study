package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	a := []int{}
	for x != 0 {
		y := x % 10
		x = x / 10
		a = append(a, y)
	}

	for i := 0; i < len(a); i++ {
		if a[i] != a[len(a)-i-1] {
			return false
		}
	}
	return true
}
