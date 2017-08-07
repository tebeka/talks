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

func producer(n int, out chan int) {
	for i := 0; i < n; i++ {
		out <- i
	}
	close(out)
}

func main() {
	in, out := make(chan int), make(chan int)
	go producer(8, in)
	go worker(7, in, out)

	for v := range out {
		fmt.Println(v)
	}
}

// END OMIT
