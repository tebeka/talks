package main

import (
	"fmt"
)

// START OMIT
type StartRequest struct {
	Image string
	Count *int
}

func main() {
	req := StartRequest{
		Image: "debian:trixie",
		Count: &3,
	}

	fmt.Println(req)
}

// END OMIT
