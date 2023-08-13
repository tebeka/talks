package main

import "fmt"

var fn = func(i int) {
	fmt.Print("Amigos")
}

func main() {
	fn := func(i int) {
		fmt.Print(i, " ")
		if i > 0 {
			fn(i - 1)
		}
	}
	fn(3)
}
