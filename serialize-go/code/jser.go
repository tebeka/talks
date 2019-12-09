package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type Payment struct {
	Name   string
	Amount float64
}

var data = `{"name": "bugs", "amount": 12.3}{"name": "daffy", "amount": 13.7}`

func main() {
	total := 0.0
	r := strings.NewReader(data)
	dec := json.NewDecoder(r)
loop:
	for {
		var p Payment
		err := dec.Decode(&p)
		switch err {
		case nil:
			total += p.Amount
		case io.EOF:
			break loop
		default:
			log.Fatal(err)
		}
	}

	fmt.Println("total:", total)
}
