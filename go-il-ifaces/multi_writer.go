package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// START_MAIN OMIT
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	log.SetOutput(io.MultiWriter(os.Stderr, file))
	log.Printf("Run Forest, Run!")
	// END_MAIN OMIT
}
