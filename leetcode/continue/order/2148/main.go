package main

import "slices"

func countElements(nums []int) int {
	m := slices.Min(nums)
	M := slices.Max(nums)

	res := 0
	for _, n := range nums {
		if n != m && n != M {
			res += 1
		}
	}

	return res
}
