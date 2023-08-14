package main

import "fmt"

type Driver struct{}

func (d *Driver) Name() string {
	return "CoolDB"
}

func main() {
	var d *Driver
	name := Driver.Name(d)
	fmt.Println(name)
}
