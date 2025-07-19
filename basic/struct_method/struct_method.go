package main

import "fmt"

type rect struct {
	width  int
	height int
}

// 指针接收
// 方法内对结构体的修改会影响到原结构体（因为接收者是原结构体的指针）。
// 如果结构体很大，使用指针接收者可以避免复制的开销（因为值接收者会在每次方法调用时复制整个结构体）。
// 对于小的结构体，使用值接收者可能更高效（因为指针操作可能涉及堆分配和垃圾回收，而小结构体的复制开销很小）。
func (r *rect) area() int {
	r.width = 100
	return r.width * r.height
}

// 值接收
// 方法内操作的是结构体的一个副本，对结构体的修改不会影响原结构体。
func (r rect) perim() int {
	r.width = 200
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("width: ", r.width)
	fmt.Println("perim:", r.perim())
	fmt.Println("width: ", r.width)

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
