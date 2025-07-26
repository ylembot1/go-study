package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 没有缓冲的channel，像一座桥，必须要等到两边都建设好，才能通行
// 有缓冲的channel，像一艘船，可以先装载货物。

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine.

func example_basic_1() {

	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	message := make(chan string)

	go func() {
		// Send a value into a channel using the ** channel <- ** syntax.
		time.Sleep(time.Second * 2)
		message <- "ping"
	}()

	// The <-channel syntax receives a value from the channel
	// By default sends and receives block until both the sender and receiver are ready.
	// This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
	fmt.Println("waiting for message...")
	msg := <-message
	fmt.Println(msg)
}

// By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value
// Buffered channels accept a limited number of values without a corresponding receiver for those values.
func example_channel_buffer() {
	message := make(chan string, 2)

	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
}

// We can use channels to synchronize execution across goroutines.
// Here’s an example of using a blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
func example_channel_Synchronization() {
	done := make(chan bool, 1)
	go func() {
		fmt.Println("working...")
		time.Sleep(time.Second * 2)
		done <- true
	}()

	// Block until we receive a notification from the worker on the channel.
	<-done
	fmt.Println("done")
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// When using channels as function parameters, you can specify if a channel is meant to only send or receive values.
// This specificity increases the type-safety of the program.
func example_channel_direction() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// Go’s select lets you wait on multiple channel operations.
// Combining goroutines and channels with select is a powerful feature of Go.
func example_channel_select() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for range 2 {
		select {
		case msg := <-c1:
			fmt.Println("received: ", msg)
		case msg := <-c2:
			fmt.Println("received: ", msg)
		}
	}
}

func example_channel_timeout() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 2")
	}
}

// Basic sends and receives on channels are blocking.
// However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.
func example_channel_non_blocking() {
	message := make(chan string)
	signals := make(chan bool)

	// Here’s a non-blocking receive.
	// If a value is available on messages then select will take the <-messages case with that value.
	// If not it will immediately take the default case.
	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// A non-blocking send works similarly.
	// Here msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver.
	// Therefore the default case is selected.
	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use multiple cases above the default clause to implement a multi-way non-blocking select.
	// Here we attempt non-blocking receives on both messages and signals.
	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

// Closing a channel indicates that no more values will be sent on it.
// This can be useful to communicate completion to the channel’s receivers.
func example_channel_close() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// Here’s the worker goroutine.
			// It repeatedly receives from jobs with j, more := <-jobs.
			// In this special 2-value form of receive, the more value will be false if jobs has been closed and all values in the channel have already been received.
			// We use this to notify on done when we’ve worked all our jobs.
			j, more := <-jobs
			if more {
				fmt.Println("received job: ", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)
	<-done

	_, more := <-jobs
	fmt.Println("more jobs? ", more)
}

func example_channel_range() {
	queue := make(chan string, 2)

	queue <- "one"
	queue <- "two"
	close(queue)

	// This range iterates over each element as it’s received from queue.
	// Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}

func example_worker_pool() {
	jobNumbers := 5
	jobs := make(chan int, jobNumbers)
	results := make(chan int, jobNumbers)

	for w := 1; w <= 3; w++ {
		go func(workerId int, jobs <-chan int, results chan<- int) {
			fmt.Println("worker initialized: ", workerId)
			for j := range jobs {
				fmt.Println("worker", workerId, "processing job", j)
				time.Sleep(time.Second)
				results <- j * 2
			}
		}(w, jobs, results)
	}

	for j := 1; j <= jobNumbers; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= jobNumbers; a++ {
		<-results
	}
}

// To wait for multiple goroutines to finish, we can use a wait group.
func example_waitgroup() {
	var wg sync.WaitGroup

	worker := func(id int) {
		fmt.Printf("Worker %d starting\n", id)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d done\n", id)
	}

	for i := 1; i <= 5; i++ {
		// Launch several goroutines and increment the WaitGroup counter for each.
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(i)
		}(i)
	}

	// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
	wg.Wait()
	fmt.Println("All workers done")
}

func example_atomic_counter() {
	var ops atomic.Int64

	var wg sync.WaitGroup

	ops.Store(1)

	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				ops.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops.Load())
}

type Counter struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[key]++
}

func example_mutex() {
	c := Counter{counters: make(map[string]int)}

	doIncrement := func(name string, n int) {
		for range n {
			c.Inc(name)
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			doIncrement("a", 1000)
		}(i)
	}

	wg.Wait()
	fmt.Println("counters:", c.counters)
}

func main() {
	example_basic_1()
	fmt.Println("--------------------------------")
	example_channel_buffer()
	fmt.Println("--------------------------------")
	example_channel_Synchronization()
	fmt.Println("--------------------------------")
	example_channel_direction()
	fmt.Println("--------------------------------")
	example_channel_select()
	fmt.Println("--------------------------------")
	example_channel_timeout()
	fmt.Println("--------------------------------")
	example_channel_non_blocking()
	fmt.Println("--------------------------------")
	example_channel_close()
	fmt.Println("--------------------------------")
	example_channel_range()
	fmt.Println("--------------------------------")
	example_worker_pool()
	fmt.Println("--------------------------------")
	example_waitgroup()
	fmt.Println("--------------------------------")
	example_atomic_counter()

	fmt.Println("--------------------------------")
	example_mutex()
}
