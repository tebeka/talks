package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("wc.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	// START OMIT
	count := 0
	s := bufio.NewScanner(file)
	for s.Scan() { // HL
		count++
	}
	if err := s.Err(); err != nil { // HL
		log.Fatalf("error: %s", err)
	}
	fmt.Println(count)
	// END OMIT
}
