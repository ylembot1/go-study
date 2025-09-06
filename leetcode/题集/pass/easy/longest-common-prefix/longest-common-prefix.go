package leetcode

func longestCommonPrefix(strs []string) string {
	var res []rune
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i || strs[j][i] != strs[0][i] {
				return string(res)
			}
		}
		res = append(res, rune(strs[0][i]))
	}
	return string(res)
}
