package main

import (
	"fmt"
	"iter"
)

// START OMIT
func Fibs() iter.Seq[int] {
	fn := func(yield func(int) bool) { // HL
		a, b := 1, 1
		for {
			if !yield(a) { // HL
				break
			}
			a, b = b, a+b
		}
	}

	return fn
}

func main() {
	for n := range Fibs() { // HL
		fmt.Println(n)
		if n > 10 {
			break
		}
	}
}

// END OMIT
