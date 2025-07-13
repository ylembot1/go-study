package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func funcVars() (int, bool) {
	return 1, true
}

// Variadic functions can be called with any number of trailing arguments.
// For example, fmt.Println is a common variadic function.
func sum(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}

	return sum
}

// -----------------------------------------
// Function closures

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println(fact(7))

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}

// func main() {
// 	res := plus(1, 2)
// 	fmt.Println("1+2 =", res)
// 	res = plusplus(1, 2, 3)
// 	fmt.Println("1+2+3 =", res)

// 	v, b := funcVars()
// 	fmt.Println("v:", v, "b:", b)

// 	fmt.Println(sum(1, 2))
// 	fmt.Println(sum(1, 2, 3))
// 	nums := []int{1, 2, 3, 4}
// 	fmt.Println(sum(nums...))
// }
