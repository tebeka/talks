package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// START_PTIME OMIT
type PTime struct {
	time.Time // HL
}

func (pt PTime) MarshalJSON() ([]byte, error) {
	// Step 1: Convert to value encoding/json can handle
	u := pt.Unix()
	// Step 2: Use json.Marshal
	return json.Marshal(u)
	// Step 3: There is no step 3
}

func (t *PTime) UnmarshalJSON(data []byte) error {
	var u int64
	if _, err := fmt.Fscanf(bytes.NewReader(data), "%d", &u); err != nil {
		return err
	}

	t.Time = time.Unix(u, 0)
	return nil
}

// END_PTIME OMIT

// START_PAYMENT OMIT
type Payment struct {
	Time   PTime // HL
	From   string
	To     string
	Amount float64
}

// END_PAYMENT OMIT

func main() {
	// START_MAIN OMIT
	data := []byte(`{
		"time": 1642777729,
		"from": "Wile. E. Coyote",
		"to":     "ACME",
		"amount": 123.45
	}`)

	var p Payment
	if err := json.Unmarshal(data, &p); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", p)
	// END_MAIN OMIT

}
