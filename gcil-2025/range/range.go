package main

import (
	"fmt"
	"iter"
)

// Range returns a sequence of integers from 0 to n-1.
func Range(n int) iter.Seq[int] {
	fn := func(yield func(int) bool) {
		for i := range n {
			if !yield(i) {
				break
			}
		}
	}

	return fn
}

func main() {
	for i := range Range(3) {
		fmt.Println(i)
	}
}
