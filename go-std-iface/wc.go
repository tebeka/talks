package main

import (
	"fmt"
	"io"
	"os"
	"strings"
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
	// START_MAIN OMIT
	poem := `
The Road goes ever on and on
Down from the door where it began.
Now far ahead the Road has gone,
And I must follow, if I can,
Pursuing it with eager feet,
Until it joins some larger way
Where many paths and errands meet.
And whither then? I cannot say.
`
	var lc LineCounter
	_, err := io.Copy(&lc, strings.NewReader(poem))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Println(lc.N)
	// END_MAIN OMIT
}
