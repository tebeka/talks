package main

import (
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	var wg sync.WaitGroup

	for i := range 4 {
		wg.Go(func() { // HL
			fmt.Println(i)
		})
	}
	wg.Wait()

	fmt.Println("Done")
	// END OMIT
}
