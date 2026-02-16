package main

import (
	"fmt"
)

// START OMIT
type StartRequest struct {
	Image string
	Count *int // HL
}

func main() {
	req := StartRequest{
		Image: "debian:trixie",
		Count: &3, // HL
	}

	fmt.Println(req)
}

// END OMIT
