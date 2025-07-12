package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If the key doesn’t exist, the zero value of the value type is returned.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	// Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, exist := m["k2"]
	fmt.Println("exist:", exist)

	n := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]int{"foo": 1, "bar": 2}
	// Equal reports whether two maps contain the same key/value pairs. Values are compared using ==.
	if maps.Equal(n, n2) {
		fmt.Println("n and n2 are equal")
	}

	// interate
	for k, v := range n {
		fmt.Println("k: ", k, "v: ", v)
	}

	elem := []string{"foo", "bar", "baz"}
	for _, v := range elem {
		if _, ok := n[v]; ok {
			fmt.Println(v, "is in the map")
		} else {
			fmt.Println(v, "is not in the map")
		}
	}

	funcMap := map[string](func(int) int){
		"int":    addOne,
		"string": plusOne,
	}

	fmt.Println(funcMap["int"](1))
	fmt.Println(funcMap["string"](len("hello")))
}

func addOne(a int) int {
	return a + 1
}

func plusOne(lenS int) int {
	return lenS + 1
}
