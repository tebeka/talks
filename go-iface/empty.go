package main

import (
	"fmt"
)

func main() {
	var i interface{}
	i = 31
	fmt.Printf("%v of type %T\n", i, i)
	i = "Bazinga"
	fmt.Printf("%v of type %T\n", i, i)

	s := i.(string)
	fmt.Println(s)

	// v := i.(int) will panic
	v, ok := i.(int)
	if ok {
		fmt.Printf("an int of value %d\n", v)
	} else {
		fmt.Println("not an int")
	}

	switch i.(type) {
	case string:
		fmt.Println("a string")
	case int:
		fmt.Println("an int")
	default:
		fmt.Printf("some other type - %T\n", i)
	}
}
