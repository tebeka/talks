package main

import (
	"fmt"
	"iter"
)

// Fibs returns a sequence of Fibonacci numbers.
func Fibs() iter.Seq[int] {
	fn := func(yield func(int) bool) {
		a, b := 0, 1
		for {
			yield(a)
			a, b = b, a+b
		}
	}

	return fn
}

func main() {
	for i := range Fibs() {
		fmt.Println(i)
		if i > 10 {
			break
		}
	}
}
