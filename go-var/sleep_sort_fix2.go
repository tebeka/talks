package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// START OMIT
	var wg sync.WaitGroup
	for _, n := range []int{3, 1, 2} {
		wg.Add(1)
		go func(n int) { // HL
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Millisecond)
			fmt.Printf("%d ", n)
		}(n) // HL
	}
	wg.Wait()
	// END OMIT
}
