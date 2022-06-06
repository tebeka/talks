package main

import (
	"fmt"
	"io"
	"os"
)

// START_COUNTER OMIT
type LineCounter struct {
	N int
}

func (lc *LineCounter) Write(data []byte) (int, error) { // HL
	for _, c := range data {
		if c == '\r' || c == '\n' {
			lc.N++
		}
	}

	return len(data), nil
}

// END_COUNTER OMIT

func main() {
	file, err := os.Open("road.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	// START_MAIN OMIT
	var lc LineCounter
	_, err = io.Copy(&lc, file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Println(lc.N)
	// END_MAIN OMIT
}
