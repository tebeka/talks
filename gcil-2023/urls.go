package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	urls := []string{
		"https://go.dev",
		"https://google.com",
		"https://httpbin.org/404",
	}

	// START OMIT
	// fan out
	ch := make(chan Result)
	for _, url := range urls {
		url := url
		go func() {
			ch <- Result{url, checkURL(url)}
		}()
	}

	// collect
	for range urls { // HL
		r := <-ch
		fmt.Printf("%s: %v\n", r.URL, r.Err)
	}
	// END OMIT

}

type Result struct {
	URL string
	Err error
}

func checkURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%q: bad status - %s", url, resp.Status)
	}

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		return fmt.Errorf("%q: can't copy - %w", url, err)
	}

	return nil
}
