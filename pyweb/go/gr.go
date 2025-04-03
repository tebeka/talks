package main

import (
	"fmt"
	"time"
)

// START_GR OMIT
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

// END_GR OMIT
