package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// START_PAYMENT OMIT
type Payment struct {
	Time   time.Time `json:"time"`
	From   string    `json:"from"`
	To     string    `json:"to"`
	Amount float64   `json:"amount"`
}

// END_PAYMENT OMIT

func demo() error {
	// START_MAIN OMIT
	data := []byte(`{
		"from": "Wile. E. Coyote",
		"to":     "ACME",
		"amount": 123.45
	}`)

	var p Payment
	if err := json.Unmarshal(data, &p); err != nil {
		return err
	}

	fmt.Printf("%#v\n", p)
	// END_MAIN OMIT
	return nil
}

func main() {
	demo()
}
