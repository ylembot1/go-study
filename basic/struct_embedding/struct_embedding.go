package main

import "fmt"

type base struct {
	num int
}

func (b *base) describe() string {
	return fmt.Sprintf("base with num: %v", b.num)
}

type container struct {
	base
	str string
}

func (c *container) describe() string {
	return fmt.Sprintf("container with num: %v and %v", c.num, c.str)
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// We can access the base’s fields directly on co, e.g. co.num
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println(co.base.describe()) // use base's describe method

	type describer interface {
		describe() string
	}

	// even though container has no describe method,
	// it implements the describer interface because it embeds base
	var d describer = &co
	fmt.Println(d.describe()) // user container's describe method
}
