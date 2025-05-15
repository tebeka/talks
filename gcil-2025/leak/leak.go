package main

import (
	"bufio"
	"fmt"
	"iter"
	"log/slog"
	"os"
	"strings"
)

// Lines return a sequence of lines file.
func Lines(fileName string) iter.Seq[string] {
	file, err := os.Open(fileName)
	if err != nil {
		slog.Warn("open", "path", fileName, "error", err)
		return func(yield func(string) bool) {
			return
		}
	}

	fn := func(yield func(string) bool) {
		s := bufio.NewScanner(file)
		for s.Scan() {
			line := s.Text()
			if !yield(line) {
				return
			}
		}
		file.Close()

		if err := s.Err(); err != nil {
			slog.Warn("scan", "path", fileName, "error", err)
			return
		}
	}

	return fn
}

func main() {
	const modFile = "go.mod"
	lnum := 0
	for line := range Lines(modFile) {
		lnum++
		if strings.HasPrefix(line, "module") {
			i := strings.Index(line, " ")
			if i < 0 {
				fmt.Fprintf(os.Stderr, "error: %s:%d - bad module line\n", modFile, lnum)
				os.Exit(1)
			}

			fmt.Println(line[i+1:])
			return
		}
	}
}
