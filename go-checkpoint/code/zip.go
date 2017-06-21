package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.checkpoint.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	out, err := os.Create("checkpoint.html.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	gz := gzip.NewWriter(out)
	defer gz.Close()
	io.Copy(gz, resp.Body)
}
