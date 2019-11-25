package main

import (
	"fmt"
	"time"
)

// Trade object
type Trade struct {
	Time   time.Time
	Symbol string
	Volume int64
	Price  float64
	IsBuy  bool `toml:"buy"`
}

func main() {
	t := Trade{
		Time:   time.Now(),
		Symbol: "BRK-A",
		Volume: 1,
		Price:  321_801.07,
		IsBuy:  true,
	}
	fmt.Printf("%#v\n", t)
	/*
		main.Trade{
			Time:time.Time{wall:0xbf6f0bebf6db79e5, ext:64591, loc:(*time.Location)(0x570b60)},
			Symbol:"BRK-A",
			Volume:1,
			Price:321801.07,
			IsBuy:true
		}
	*/
}
