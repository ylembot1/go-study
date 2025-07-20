package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

var ErrOther = fmt.Errorf("other error")
var ErrOutOfTea = fmt.Errorf("no more tea available: %w", ErrOther)
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		// We can wrap errors with higher-level errors to add context.
		// The simplest way to do this is with the %w verb in fmt.Errorf.
		// Wrapped errors create a logical chain (A wraps B, which wraps C, etc.) that can be queried with functions like errors.Is and errors.As.
		return fmt.Errorf("can't boil water: %w", ErrPower)
	}

	return nil
}

type argError struct {
	arg int
	msg string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.msg)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with 42"}
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("out of tea")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("can't boil water")
			} else {
				fmt.Println("unknown error")
			}
		}
	}

	_, err := f2(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.msg)
	} else {
		fmt.Println("err doesn't match argError")
	}

	fmt.Println("Tea is ready!")

}
