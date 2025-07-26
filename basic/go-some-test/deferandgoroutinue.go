package main

import (
	"fmt"
	"time"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %.2f seconds\n", name, elapsed.Seconds())
}

func foo() {

	go func() {
		defer TimeTrack(time.Now(), "foo inner")
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine done")
	}()

}

func main() {
	foo()
	time.Sleep(time.Second * 3)
	fmt.Println("main done")
}
