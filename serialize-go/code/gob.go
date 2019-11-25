package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

// Trade object
type Trade struct {
	Time   time.Time
	Symbol string
	Volume int64
	Price  float64
	IsBuy  bool
}

func main() {
	t1 := Trade{
		Time:   time.Now(),
		Symbol: "BRK-A",
		Volume: 1,
		Price:  321_801.07,
		IsBuy:  true,
	}

	var network bytes.Buffer
	gob.NewEncoder(&network).Encode(t1)
	var t2 Trade
	gob.NewDecoder(&network).Decode(&t2)
	fmt.Printf("%+v\n", t2) // {Time:2019-11-25 12:41:04.41768293 +0200 IST Symbol:BRK-A Volume:1 Price:321801.07 IsBuy:true}
}
