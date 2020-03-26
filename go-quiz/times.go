package main

import (
	"fmt"
	"time"
)

type Log struct {
	Message string
	time.Time
}

func main() {
	ts, err := time.Parse("2006-01-02", "2020-03-21")
	if err != nil {
		panic(err)
	}
	log := Log{"Hello", ts}
	fmt.Printf("%v\n", log)
}
