package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
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
	var data = `{"type":"click","x":100,"y":200}{"type":"move","x":101,"y":202}`
	dec := json.NewDecoder(strings.NewReader(data))

	for {
		var e Event
		err := dec.Decode(&e)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		fmt.Println(e)
	}
	return nil
	// END_MAIN OMIT
}

func main() {
	demo()
}
