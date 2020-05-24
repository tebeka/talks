package main

import (
	"os"
	"time"

	"gopkg.in/yaml.v2"
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
	t1 := Trade{
		Time:   time.Now(),
		Symbol: "BRK-A",
		Volume: 1,
		Price:  321_801.07,
		IsBuy:  true,
	}

	yaml.NewEncoder(os.Stdout).Encode(t1)
}
