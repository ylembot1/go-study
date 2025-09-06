package leetcode

import "slices"

func maxScoreIndices(nums []int) []int {
	n1 := 0
	for _, i := range nums {
		n1 += i
	}
	// n0 := len(nums) - n1

	s := make([]int, len(nums)+1)
	ls := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			ls++
		} else {
			n1--
		}
		s[i+1] = ls + n1
	}
	s[0] = n1

	res := make([]int, 0)
	max := slices.Max(s)
	for idx, v := range s {
		if v == max {
			res = append(res, idx)
		}
	}

	return res
}
