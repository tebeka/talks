package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d ", i)
			time.Sleep(time.Second)
		}
	}()
	runtime.Goexit()
}
