package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	b := s[:]
	s = append(s, 4)
	s[1] = 99
	fmt.Println(b)
}
