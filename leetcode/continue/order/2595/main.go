package main

func evenOddBit(n int) []int {
	c := 0
	ans := []int{0, 0}

	for n != 0 {
		if n&1 == 1 {
			ans[c%2]++
		}
		c++
		n >>= 1
	}

	return ans
}
