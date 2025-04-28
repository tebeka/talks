package main

import (
	"fmt"
	"iter"
	"time"
)

func Tick(d time.Duration) iter.Seq[time.Time] {
	fn := func(yield func(time.Time) bool) {
		for {
			if !yield(time.Now()) {
				return
			}
			time.Sleep(d)
		}
	}

	return fn
}

func main() {
	i := 0
	for t := range Tick(time.Second) {
		fmt.Println(t)
		i++
		if i > 5 {
			break
		}
	}
}
