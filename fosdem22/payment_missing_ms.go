package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mitchellh/mapstructure"
)

// START_PAYMENT OMIT
type Payment struct {
	Time   time.Time `json:"time"`
	From   string    `json:"from"`
	To     string    `json:"to"`
	Amount float64   `json:"amount"`
}

// END_PAYMENT OMIT

func main() {
	// START_MAIN OMIT
	data := []byte(`{
		"from": "Wile. E. Coyote",
		"to":     "ACME",
		"amount": 123.45
	}`)

	var m map[string]interface{} // HL
	if err := json.Unmarshal(data, &m); err != nil {
		log.Fatal(err)
	}

	var p Payment
	if err := mapstructure.Decode(m, &p); err != nil { // HL
		log.Fatal(err)
	}

	if _, ok := m["time"]; !ok { // HL
		p.Time = time.Now()
	}
	fmt.Printf("%#v\n", p)
	// END_MAIN OMIT
}
