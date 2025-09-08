package main

import "math"

func isThree(n int) bool {

	if n < 4 {
		return false
	}

	t := int(math.Sqrt(float64(n)))
	for i := 2; i < t; i++ {
		if n%i == 0 {
			return false
		}
	}
	return t*t == n
}
