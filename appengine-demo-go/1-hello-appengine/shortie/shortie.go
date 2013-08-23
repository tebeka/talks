/* AppEngine demo - a URL shortener */
//package shortie
package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", rootHandler)
}

// rootHandler handles the main page.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello AppEngine\n")
}

func main() {
	http.ListenAndServe(":8080", nil)
}
