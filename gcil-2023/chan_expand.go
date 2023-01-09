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
		close(ch)
	}()

	for { // HL
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	// END OMIT
}
