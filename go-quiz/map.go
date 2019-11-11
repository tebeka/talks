package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["errors"]++
	fmt.Println(m["errors"])
}
