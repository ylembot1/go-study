package main

func pivotInteger(n int) int {
	mp := map[int]int{}

	t := 0
	for i := range n + 1 {
		t += i
		mp[t] = i

	}

	for k, v := range mp {
		if t-k+v == k {
			return v
		}
	}

	return -1
}
