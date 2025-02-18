package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("%d ", v)
	}
}
