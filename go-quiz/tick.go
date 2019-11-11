package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(100)
	for range t.C {
		t.Stop()
		fmt.Printf("bye ")
	}
	fmt.Println("bye")
}
