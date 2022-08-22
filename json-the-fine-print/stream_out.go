package main

import (
	"encoding/json"
	"os"
)

// START_EVENT OMIT
type Event struct {
	Type string  `json:"type"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
}

// END_EVENT OMIT

func demo() error {
	// START_MAIN OMIT
	events := []Event{
		{"click", 100, 200},
		{"move", 101, 202},
	}
	enc := json.NewEncoder(os.Stdout)
	for _, e := range events {
		if err := enc.Encode(e); err != nil {
			return err
		}
	}
	// END_MAIN OMIT
	return nil
}

func main() {
	demo()
}
