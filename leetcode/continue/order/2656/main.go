func maximizeSum(nums []int, k int) int {
	M := slices.Max(nums)
	res := 0
	for k > 0 {
		k -= 1
		res += M + k
	}

	return res
}