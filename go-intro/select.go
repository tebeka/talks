package main

import (
	"fmt"
	"time"
)

// START OMIT
func worker(ch chan int, n int, timeout time.Duration) {
	time.Sleep(timeout)
	ch <- n
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go worker(ch1, 1, time.Millisecond*100)
	go worker(ch2, 2, time.Millisecond*200)

	select {
	case v := <-ch1:
		fmt.Printf("%d from ch1\n", v)
	case v := <-ch2:
		fmt.Printf("%d from ch2\n", v)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}

// END OMIT
