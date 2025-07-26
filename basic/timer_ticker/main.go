package main

import (
	"fmt"
	"time"
)

func timers() {
	// Timers represent a single event in the future.
	// You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.
	// This timer will wait 2 seconds.
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	timer2.Reset(time.Second * 1)
	select {
	case <-timer2.C:
		fmt.Println("Timer 2 fired again")
	case <-time.After(time.Second * 2):
		fmt.Println("Timer 2 not fired")
	}

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

func tickers_use_range() {
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func tickers_use_select() {
	ticker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

// We often want to execute Go code at some point in the future, or repeatedly at some interval.
// Go’s built-in timer and ticker features make both of these tasks easy.

// Timers are for when you want to do something once in the future
// Tickers are for when you want to do something repeatedly at regular intervals.

func main() {

	timers()

	fmt.Println("--------------------------------")

	tickers_use_range()

	fmt.Println("--------------------------------")

	tickers_use_select()

}
