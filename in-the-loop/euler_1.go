// https://projecteuler.net/problem=1
package main

import (
	"fmt"
)

func main() {

	total := 0
	for i := 0; i < 1_000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	fmt.Println(total)
}
