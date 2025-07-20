package main

import (
	"cmp"
	"fmt"
	"slices"
)

func builtinSort() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("strs:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("ints:", ints)

	// IsSorted reports whether x is sorted in ascending order
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}

func sortByFunc() {
	fruits := []string{"peach", "banana", "kiwi", "apple"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println("fruits:", fruits)

	type peopleStruct struct {
		name string
		age  int
	}

	people := []peopleStruct{
		{name: "Jax", age: 37},
		{name: "Lena", age: 29},
		{name: "Mike", age: 32},
		{name: "Alice", age: 29},
		{name: "Bob", age: 32},
	}

	peopleCmp := func(a, b peopleStruct) int {
		if a.age == b.age {
			return cmp.Compare(b.name, a.name) // 名字降序
		}

		return b.age - a.age // 降序  // 等价cmp.Compare(b.age, a.age)
		// return cmp.Compare(a.age, b.age)  // 等价与a.age - b.age  // 升序
	}

	slices.SortFunc(people, peopleCmp)
	fmt.Println("people:", people)
}

func main() {
	builtinSort()

	sortByFunc()
}
