package main

import "fmt"

// Go’s structs are typed **collections of fields**.
// They’re useful for grouping data together to form records.
type Person struct {
	name string
	age  int
}

// Go is a garbage collected language;
// you can safely return a pointer to a local variable
// - it will only be cleaned up by the garbage collector when there are no active references to it.
func newPerson(name string) *Person {
	p := new(Person)
	p.name = name
	p.age = 0
	return p

	// p := person{name: name}
	// p.age = 42
	// return &p
}

func main() {
	fmt.Println(Person{"Bob", 20})

	fmt.Println(Person{name: "Alice", age: 30})

	fmt.Println(Person{name: "Fred"})

	fmt.Println(&Person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	// If a struct type is only used for a single value, we don’t have to give it a name.
	// The value can have an anonymous struct type.
	// This technique is commonly used for table-driven tests.

	dogs := []struct {
		name   string
		isGood bool
	}{
		{
			name:   "Rex",
			isGood: true,
		},
		{
			name:   "Fido",
			isGood: false,
		},
	}
	fmt.Println(dogs)

}
