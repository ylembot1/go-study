package leetcode

func checkString(s string) bool {
	hasB := false

	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			hasB = true
		} else if s[i] == 'a' && hasB {
			return false
		}
	}
	return true
}
