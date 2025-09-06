package leetcode

import "strings"

func lengthOfLastWord1(s string) int {
	trimStr := strings.TrimSpace(s)

	splitedStr := strings.Fields(trimStr)

	if len(splitedStr) <= 0 {
		return 0
	}
	return len(splitedStr[len(splitedStr)-1])
}

func lengthOfLastWord(s string) int {
	res := 0
	tmp := 0
	s += " "
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if tmp != 0 {
				res = tmp
				tmp = 0
			}
		} else {
			tmp += 1
		}
	}
	return res
}
