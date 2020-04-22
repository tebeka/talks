// https://golang.org/doc/effective_go.html#interfaces
package main

import (
	"fmt"
	"log"
	"net/http"
)

type Counter int

func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	*c++
	fmt.Fprintf(w, "count = %d\n", *c)
}

func main() {
	var c Counter
	http.Handle("/count", &c)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
