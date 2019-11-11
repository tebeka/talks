package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	b := s[:]
	s[1] = 99
	fmt.Println(b)
}
