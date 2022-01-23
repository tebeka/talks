package main

import (
	"fmt"
	"io"
	"os"
)

// START_RANGE OMIT
func ReadRange(r io.Reader, start int64, buf []byte) (int, error) {
	if ra, ok := r.(io.ReaderAt); ok { // HL
		return ra.ReadAt(buf, start)
	}

	if s, ok := r.(io.Seeker); ok { // HL
		if _, err := s.Seek(start, io.SeekStart); err != nil {
			return 0, err
		}
	} else { // Slow seek
		if _, err := io.Copy(io.Discard, io.LimitReader(r, start)); err != nil {
			return 0, err
		}
	}

	return r.Read(buf)
}

// END_RANGE OMIT

func main() {
	file, err := os.Open("road.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	// START_MAIN OMIT
	buf := make([]byte, 10)
	if _, err := ReadRange(file, 7, buf); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Println(string(buf))
	// END_MAIN OMIT
}
