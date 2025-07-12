package main

import "fmt"

func main() {

	// Here we create an array a that will hold **exactly** 5 ints.
	// The **type of elements and **length** are both part of the array’s type.
	// By default an array is zero-valued, which for ints means 0s.
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// The builtin len returns the length of an array.
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3}
	fmt.Println("dcl: ", b)

	// b = [...]int{1, 2, 3} // This is not allowed. because the length of the array is fixed.
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl...: ", b)
	fmt.Println("len: ", len(b))

	// If you specify the index with :, the elements in between will be zeroed.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var c [2][3]int
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("2d: ", c)
}
