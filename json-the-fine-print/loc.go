package main

import (
	"encoding/json"
	"fmt"
)

// START_LOC OMIT
type Location struct {
	Lat float64
	Lng float64
}

// END_LOC OMIT

func demo() error {
	// START_MAIN OMIT
	data := []byte(`{"lat": 132.5270941, "lng": 34.9404309}`) // HL

	var loc Location
	if err := json.Unmarshal(data, &loc); err != nil {
		return err
	}

	fmt.Println(loc)
	// END_MAIN OMIT
	return nil
}

func main() {
	demo()
}
