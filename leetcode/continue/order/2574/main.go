func leftRightDifference(nums []int) []int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	// right := make([]int, len(nums))

	// i := len(nums) - 2
	// for i >= 0 {
	//     right[i] = right[i+1] + nums[i+1]
	//     i -= 1
	// }

	right := 0
	for _, x := range nums {
		right += x
	}

	ans := make([]int, len(nums))
	left := 0
	for i, n := range nums {
		right -= n
		ans[i] = abs(left - right)
		left += n

		// ans[i] = abs(left - right[i])
		// left += n
	}

	return ans
}