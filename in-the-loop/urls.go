package main

import (
	"fmt"
)

func checkURL(url string) error {
	return nil
}

type Result struct {
	URL string
	Err error
}

func main() {
	urls := []string{
		"https://google.com",
		"https://apple.com",
		"https://ibm.com",
	}
	// fan out
	ch := make(chan Result)
	for _, url := range urls {
		url := url
		go func() {
			ch <- Result{url, checkURL(url)}
		}()
	}
	// collect
	for range urls {
		r := <-ch
		fmt.Printf("%s: %v\n", r.URL, r.Err)
	}
}
