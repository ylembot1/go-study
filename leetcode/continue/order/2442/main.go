func countDistinctIntegers(nums []int) int {
	fn := func(x int) int {
		res := 0
		for x != 0 {
			res = res*10 + x%10
			x /= 10
		}

		return res
	}

	mp := map[int]bool{}
	for _, n := range nums {
		mp[n] = true
		mp[fn(n)] = true
	}

	return len(mp)
}