package main

import (
	"encoding/json"
	"log"
	"os"
)

// START_EVENT OMIT
type Event struct {
	Type string  `json:"type"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
}

// END_EVENT OMIT

func main() {
	// START_MAIN OMIT
	events := []Event{
		{"click", 100, 200},
		{"move", 101, 202},
	}
	enc := json.NewEncoder(os.Stdout)
	for _, e := range events {
		if err := enc.Encode(e); err != nil {
			log.Fatal(err)
		}
	}
	// END_MAIN OMIT
}
