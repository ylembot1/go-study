package main

import (
	"fmt"
	"slices"
)

func main() {
	// Unlike arrays, **slices are typed only by the elements they contain** (not the number of elements).
	// An uninitialized slice equals to nil and has length 0.
	s := []string{}
	fmt.Println("uninit:", s, "len:", len(s), "cap:", cap(s))

	// To create a slice with non-zero length, use the builtin make.
	// Here we make a slice of strings of length 3 (initially zero-valued).
	// By default a new slice’s capacity is equal to its length;
	// if we know the slice is going to grow ahead of time, it’s possible to pass a capacity explicitly as an additional parameter to make.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices support several more that make them richer than arrays.
	// One is the builtin append, which returns a slice containing one or more new values.
	// Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	// Slices can also be copy’d.
	// Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s) // deep copy
	c[0] = "x"
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	// We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Equal reports whether two slices are equal: the same length and all elements equal
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal")
	}

	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		twoD[i] = make([]int, 2)
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	arr := [5]int{1, 2, 3, 4, 5}
	s3 := arr[1:4] // len=3, cap=4，共享arr的内存
	fmt.Println("s3:", s3, "len:", len(s3), "cap:", cap(s3))
	s3 = append(s3, 6)
	fmt.Println("s3:", s3, "len:", len(s3), "cap:", cap(s3))
	fmt.Println("arr:", arr)
	s3 = append(s3, 7)
	fmt.Println("s3:", s3, "len:", len(s3), "cap:", cap(s3))
	fmt.Println("arr:", arr)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s1 = append(s1, s2...)
	fmt.Println("s1:", s1)

	arr1 := [4]int{10, 20, 30, 40}
	slice := arr1[0:2]
	testSlice1 := slice
	testSlice2 := append(append(append(slice, 1), 2), 3)
	slice[0] = 11

	fmt.Println(testSlice1[0]) // 11
	fmt.Println(testSlice2[0]) // 10
	fmt.Println(arr1)          // [11 20 1 2]
	fmt.Println("len:", len(slice), "cap:", cap(slice))
	fmt.Println("len:", len(testSlice1), "cap:", cap(testSlice1))
	fmt.Println("len:", len(testSlice2), "cap:", cap(testSlice2))
}
