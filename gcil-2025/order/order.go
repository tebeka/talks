// Return first 1000 logs that are errors (status >= 400)
package main

import (
	"fmt"
	"iter"
	"mordor/lazy"
	"mordor/log"
	"os"
	"slices"
)

// Head returns a sequence of up to "n" items from seq.
func Head[T any](seq iter.Seq[T], n int) iter.Seq[T] {
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
	logs, error := lazy.LoadLogs("logs.json.gz")
	if error != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", error)
		os.Exit(1)
	}

	logs = Head(logs, 1000)
	logs = lazy.Filter(logs, isError)

	fmt.Println(len(slices.Collect(logs)))
}
