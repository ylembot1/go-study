package main

import "fmt"

func countSubarrays(nums []int) int {
	judge := func(end int) bool {
		return (nums[end]+nums[end-2])*2 == nums[end-1]
	}

	res := 0
	for idx, _ := range nums {
		if idx < 2 {
			continue
		}

		if judge(idx) {
			res++
		}
	}

	return res
}

func main() {
	nums := []int{1, 2, 1, 4, 1}
	fmt.Println(countSubarrays(nums))
}
