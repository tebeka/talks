package main

import "fmt"

type StartVM struct {
	Image string
	Count *int
}

func main() {
	req := StartVM{
		Image: "debian:trixie",
		Count: &3,
	}

	fmt.Println(req.Image, *req.Count)
}
