package main

import "fmt"

func main() {
	var m map[string]int
	m["errors"]++
	fmt.Println(m["errors"])
}
