package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("fgrep.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	var r io.Reader = file
	term := "err"

	// START OMIT
	lnum := 0
	s := bufio.NewScanner(r)
	for s.Scan() { // HL
		lnum++
		if strings.Contains(s.Text(), term) {
			fmt.Printf("%d: %s\n", lnum, s.Text())
		}
	}
	if err := s.Err(); err != nil { // HL
		log.Fatalf("error: %s", err)
	}
	// END OMIT
}
