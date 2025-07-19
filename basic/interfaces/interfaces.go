package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r *rect) perim() float64 {
	return 2 * (r.width + r.height)
}

var _ geometry = (*rect)(nil)

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

var _ geometry = (*circle)(nil)

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectType(g geometry) {
	switch g.(type) {
	case *rect:
		fmt.Println("rect")
	case *circle:
		fmt.Println("circle")
	default:
		fmt.Println("unknown")
	}
}

func detectCircle(g geometry) {
	if c, ok := g.(*circle); ok {
		fmt.Println("circle radius: ", c.radius)
	}
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(&r)
	measure(&c)

	detectType(&r)
	detectType(&c)

	detectCircle(&r)
	detectCircle(&c)

}
