package main

import (
	"iter"
	"mordor/log"
)

// Limit returns a sequence of up to "n" items from seq.
func Limit[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	fn := func(yield func(T) bool) {
		i := 0
		for v := range seq {
			if !yield(v) {
				return
			}

			i++
			if i == n {
				return
			}
		}
	}

	return fn
}

func isError(log log.Log) bool {
	return log.Status >= 400
}

func main() {
}
