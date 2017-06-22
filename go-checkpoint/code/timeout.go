package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetch(url string, out chan *http.Response) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		close(out)
		return
	}
	out <- resp
}

func main() {
	url := "https://www.checkpoint.com"
	out := make(chan *http.Response)

	go fetch(url, out)

	select {
	case resp, ok := <-out:
		if !ok {
			break
		}
		fmt.Printf("got %d from %s\n", resp.StatusCode, url)
	case <-time.After(300 * time.Millisecond):
		fmt.Printf("timeout")
	}
}
