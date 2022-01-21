package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// BEGIN_HANDLER OMIT
func handler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher) // HL
	if !ok {
		http.Error(w, "no streaming support", http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	for i := 0; i < 10; i++ {
		if err := enc.Encode(Point{i, i}); err != nil {
			// Can't set error
			log.Printf("encoding: %s", err)
			return
		}
		flusher.Flush() // HL
	}
}

// END_HANDLER OMIT

func main() {
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
