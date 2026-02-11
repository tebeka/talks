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
		Count: new(3), // HL
	}

	fmt.Println(req)
}

// END OMIT
