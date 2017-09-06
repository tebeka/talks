package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// START_MAIN OMIT
	file, err := os.Open("hash.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := sha256.New()
	io.Copy(hash, file)
	fmt.Printf("%x\n", hash.Sum(nil))
	// END_MAIN OMIT
}
