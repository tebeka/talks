package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// START_LOC OMIT
type Location struct {
	Lat float64
	Lng float64
}

// END_LOC OMIT

// START_VALIDATE OMIT
func (loc Location) Validate() error {
	if loc.Lat < -90 || loc.Lat > 90 {
		return fmt.Errorf("invalid latitude: %#v\n", loc.Lat)
	}

	if loc.Lng < -180 || loc.Lng > 180 {
		return fmt.Errorf("invalid longitude: %#v\n", loc.Lng)
	}

	return nil
}

// END_VALIDATE OMIT

func main() {
	// START_MAIN OMIT
	data := []byte(`{"lat": 132.5270941, "lng": 34.9404309}`)

	var loc Location
	if err := json.Unmarshal(data, &loc); err != nil {
		log.Fatal(err)
	}

	if err := loc.Validate(); err != nil { // HL
		log.Fatal(err)
	}

	fmt.Println(loc)
	// END_MAIN OMIT
}
