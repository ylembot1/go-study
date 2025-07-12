package main

import "fmt"

// range iterates over elements in a variety of built-in data structures.
func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	for idx := range nums {
		if idx == 2 {
			fmt.Println("idx: ", idx, "num: ", nums[idx])
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over just the keys of a map
	for k := range kvs {
		fmt.Println("key:", k)
		fmt.Println("value: ", kvs[k])
	}

	// range on strings iterates over Unicode code points.
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	for n := range 10 {
		fmt.Println(n)
	}

}
