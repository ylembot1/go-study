package main

import (
	"fmt"
)

// panic:
// A panic typically means something went unexpectedly wrong.
// Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.
// Note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible.

// recover
// Go makes it possible to recover from a panic, by using the recover built-in function.
// A recover can stop a panic from aborting the program and let it continue with execution instead.

// An example of where this can be useful: a server wouldn’t want to crash if one of the client connections exhibits a critical error.
// Instead, the server would want to close that connection and continue serving other clients.
// In fact, this is what Go’s net/http does by default for HTTP servers.

func mayPanic() {
	panic("a problem")
}

func main() {

	// recover must be called within a deferred function.
	// **When the enclosing function panics, the defer will activate and a recover call within it will catch the panic.**
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	fmt.Println("Calling mayPanic()")

	mayPanic()

	// This code will not run, because mayPanic panics.
	// The execution of main stops at the point of the panic and resumes in the deferred closure.
	fmt.Println("After mayPanic()")
}
