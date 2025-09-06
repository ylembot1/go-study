package leetcode

// func plusOne(digits []int) []int {
// 	v := 0
// 	for i := 0; i < len(digits); i++ {
// 		v = v*10 + digits[i]
// 	}
// 	v += 1

// 	res := make([]int, 0)
// 	for v != 0 {
// 		yu := v % 10
// 		v /= 10
// 		res = append(res, yu)
// 	}

// 	for i, j := 0, len(res)-1; i < j; {
// 		t := res[i]
// 		res[i] = res[j]
// 		res[j] = t
// 		i++
// 		j--
// 	}

// 	return res
// }

func plusOne(digits []int) []int {
	add := 1
	for i := len(digits) - 1; i >= 0 && add > 0; i-- {
		t := digits[i] + add
		add = t / 10
		digits[i] = t % 10
	}

	if add > 0 {
		return append([]int{add}, digits...)
	}
	return digits
}
