package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := 3
	fmt.Printf("before ")
	time.Sleep(timeout * time.Millisecond)
	fmt.Println("after")
}
