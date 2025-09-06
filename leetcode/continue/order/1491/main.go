package main

import "sort"

func average(salary []int) float64 {
	sort.Ints(salary)

	total := 0.0
	for i, num := range salary {
		if i == 0 || i == len(salary)-1 {
			continue
		}
		total += float64(num)
	}

	return total / float64(len(salary)-2)
}

// func average(salary []int) float64 {
//     s := 0
//     for _, x := range salary {
//         s += x
//     }
//     m := slices.Min(salary)
//     M := slices.Max(salary)
//     n := len(salary)
//     return float64(s-m-M) / float64(n-2)
// }
