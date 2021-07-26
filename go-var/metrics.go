package main

import (
	"expvar"
	"fmt"
	"log"

	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello RedHat!\n")
}

// addMetrics is a middleware
// START OMIT
func addMetrics(name string, handler http.Handler) http.Handler {
	name = fmt.Sprintf("%s.calls", name)
	nCalls := expvar.NewInt(name) // HL

	fn := func(w http.ResponseWriter, r *http.Request) {
		nCalls.Add(1) // HL
		handler.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// END OMIT

func main() {
	// START_MAIN OMIT
	handler := addMetrics("hello", http.HandlerFunc(helloHandler))
	http.Handle("/hello", handler)
	// END_MAIN OMIT
	//http.HandleFunc("/hello", helloHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
