package main

import (
	"fmt"
)

func main() {
	// START OMIT
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch) // HL
	}()

	for v := range ch { // HL
		fmt.Println(v)
	}
	// END OMIT
}
