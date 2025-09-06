package leetcode

func findKOr(nums []int, k int) int {
	t := 1
	res := 0
	for i := 1; i <= 32; i++ {
		c := 0
		for _, n := range nums {
			if n&t != 0 {
				c++
			}
		}
		if c >= k {
			res |= t
		}
		t <<= 1
	}

	return res
}
