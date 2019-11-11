package main

import "fmt"

func resolve(n, dfault int) {
	if n == 0 {
		return default
	}
	return n
}

func main() {
	v := resolve(0, 3)
	fmt.Println(v)
}
