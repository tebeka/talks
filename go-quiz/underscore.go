package main

import "fmt"

func main() {
	s := struct{ a, _, c int }{1, 2, 3}
	fmt.Println(s)
}
