package leetcode

func strStr(haystack string, needle string) int {
	i := 0
	for i < len(haystack) {
		ti := i
		j := 0
		for ; i < len(haystack) && j < len(needle); j++ {
			if haystack[i] != needle[j] {
				break
			}
			i++
		}
		if j >= len(needle) {
			return ti
		} else {
			i = ti + 1
		}
	}
	return -1
}
