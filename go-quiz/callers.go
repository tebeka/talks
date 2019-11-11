package main

import "fmt"

type stack []uintptr

func callers() stack {
	return make([]uintptr, 3)
}

func main() {
	c := callers()
	fmt.Println(c)
}
