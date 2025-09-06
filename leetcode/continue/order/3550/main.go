package main

import "fmt"

// func smallestIndex(nums []int) int {
//     calc := func (a int) int {
//         res := 0
//         for a != 0 {
//             res = res + a % 10
//             a = a / 10
//         }
//         return res
//     }

//     for idx, num := range nums {
//         res := calc(num)
//         if idx == res {
//             return idx
//         }
//     }

//     return -1
// }

func smallestIndex(nums []int) int {
	digitSum := func(n int) int {
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		return sum
	}

	for idx, num := range nums {
		if idx == digitSum(num) {
			return idx
		}
	}

	return -1
}

func main() {
	nums := []int{1, 10, 2}
	fmt.Println(smallestIndex(nums))
}
