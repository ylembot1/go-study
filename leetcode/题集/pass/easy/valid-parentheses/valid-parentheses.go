package leetcode

func isValid(s string) bool {
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	st := make([]rune, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(st) == 0 || m[rune(st[len(st)-1])] != rune(s[i]) {
				return false
			} else {
				st = st[:len(st)-1]
			}
		} else {
			st = append(st, rune(s[i]))
		}
	}
	return len(st) == 0
}
