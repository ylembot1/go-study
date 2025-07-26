package main

import (
	"fmt"
	"time"
)

func f(from string) {
	fmt.Println("from: ", from)
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println("from: ", msg)
	}("going")

	time.Sleep(time.Second * 2)
}
