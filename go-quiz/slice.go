package main

import "fmt"

func main() {
	cart := []string{"apple", "banana", "milk"}
	fruit := cart[:2]
	cart[0] = "pear"
	fmt.Println(fruit)
}
