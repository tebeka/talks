package main

import (
	"fmt"
)

// START OMIT
func worker(n int, in, out chan int) {
	for v := range in {
		out <- v * n
	}
	close(out)
}

func producer(n int, in chan int) {
	for i := 0; i < n; i++ {
		in <- i
	}
	close(in)
}

func main() {
	in, out := make(chan int), make(chan int)
	go producer(3, in)
	go worker(7, in, out)

	for v := range out {
		fmt.Println(v)
	}
}

// END OMIT
