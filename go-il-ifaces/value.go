package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// START_VALUE OMIT
type Unit string

const (
	Celcius    Unit = "c"
	Fahrenheit Unit = "f"
	Kelvin     Unit = "k"
)

// Value if a measurement value
type Value struct {
	Amount float64 `json:"value"`
	Unit   Unit    `json:"unit"`
}

// END_VALUE OMIT

// START_MJ OMIT
func (v Value) MarshalJSON() ([]byte, error) {
	// Step 1: Convert to a type encoding/json can handle
	s := fmt.Sprintf("%f%s", v.Amount, v.Unit)
	// Step 2: Use json.Marshal
	return json.Marshal(s)
	// Step 3: There is no step 3
}

// END_MJ OMIT

func main() {
	// START_MAIN OMIT
	v := Value{15.4, Celcius}
	json.NewEncoder(os.Stdout).Encode(v)
	// END_MAIN OMIT
}
