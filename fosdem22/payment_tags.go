package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// START_PAYMENT OMIT
type Payment struct {
	Time   time.Time `json:"time"`
	From   string    `json:"from"`
	To     string    `json:"to"`
	Amount float64   `json:"amount"`
}

func main() {
	p := Payment{
		Time:   time.Now(),
		From:   "Wile. E. Coyote",
		To:     "", // HL
		Amount: 123.45,
	}
	data, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
}

// END_PAYMENT OMIT
