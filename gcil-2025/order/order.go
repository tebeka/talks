// Return first 1000 logs that are errors (status >= 400)
package main

import (
	"compress/gzip"
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
	const fileName = "logs.json.gz"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer gz.Close()

	logs := lazy.LoadLogs(gz)

	logs = Head(logs, 1000)
	logs = lazy.Filter(logs, isError)

	count := len(slices.Collect(logs))
	fmt.Println(count)
}
