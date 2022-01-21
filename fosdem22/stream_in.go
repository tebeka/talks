package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
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
	var data = `{"type":"click","x":100,"y":200}{"type":"move","x":101,"y":202}`
	dec := json.NewDecoder(strings.NewReader(data))
	for {
		var e Event
		err := dec.Decode(&e)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(e)
	}
	// END_MAIN OMIT
}
