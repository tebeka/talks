package main

import (
	"fmt"
	"iter"
)

func Fixture() iter.Seq[bool] {
	fn := func(yield func(bool) bool) {
		fmt.Println(">>> setup")
		if !yield(true) {
			return
		}

		fmt.Println(">>> teardown")
		if !yield(false) {
			return
		}
	}

	return fn
}

func run(fixture iter.Seq[bool]) {
	next, stop := iter.Pull(fixture)
	defer stop()

	fmt.Println("calling setup")
	next()

	fmt.Println("running test")

	fmt.Println("calling teardown")
	next()
}

func main() {
	run(Fixture())
}
