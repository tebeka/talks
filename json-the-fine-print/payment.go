package main

import (
	"encoding/json"
	"os"
	"time"
)

// START_PAYMENT OMIT
type Payment struct {
	Time   time.Time // HL
	From   string
	To     string
	Amount float64
}

// END_PAYMENT OMIT

func main() {
	// START_MAIN OMIT
	p := Payment{
		Time:   time.Now(),
		From:   "Wile. E. Coyote",
		To:     "ACME",
		Amount: 123.45,
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(p)
	// END_MAIN OMIT
}
