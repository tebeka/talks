package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

// START_SHA OMIT
func fileSHA1(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r io.Reader = file
	if strings.HasSuffix(fileName, ".gz") { // HL
		var err error
		r, err = gzip.NewReader(file) // HL
		if err != nil {
			return "", err
		}
	}

	h := sha1.New()                          // HL
	if _, err := io.Copy(h, r); err != nil { // HL
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// END_SHA OMIT

func main() {
	sig, err := fileSHA1("system.log.gz")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Println(sig)

}
