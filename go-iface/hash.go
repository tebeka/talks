package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net/http"
)

func urlSha1(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var r io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		var err error
		r, err = gzip.NewReader(r)
		if err != nil {
			return "", err
		}

	}

	h := sha1.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func main() {
	// hash, err := urlSha1("http://httpbin.org/gzip")
	hash, err := urlSha1("http://localhost:9999/gzip")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(hash)
}
