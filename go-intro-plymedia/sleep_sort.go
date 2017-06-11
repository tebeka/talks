package main

import (
	"fmt"
	"time"
)

// START_SS OMIT
func sleepSort(nums []int) {
	for _, n := range nums {
		go func(v int) {
			time.Sleep(time.Duration(v) * time.Millisecond)
			fmt.Println(v)
		}(n)
	}
}

func main() {
	sleepSort([]int{23, 42, 4, 15, 16, 8})
	time.Sleep(time.Second)
}

// END_SS OMIT
