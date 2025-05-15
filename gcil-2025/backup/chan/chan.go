package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// Lines return a channel that reads lines from a file.
func Lines(fileName string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		file, err := os.Open(fileName)
		if err != nil {
			slog.Warn("open file", "error", err)
			return
		}
		defer file.Close()

		s := bufio.NewScanner(file)
		for s.Scan() {
			out <- s.Text()
		}
		if err := s.Err(); err != nil {
			slog.Warn("scan file", "error", err)
			return
		}
	}()

	return out
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
