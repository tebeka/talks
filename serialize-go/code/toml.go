package main

import (
	"os"
	"time"

	"github.com/BurntSushi/toml"
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

	toml.NewEncoder(os.Stdout).Encode(t1)
}
