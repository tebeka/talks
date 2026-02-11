package main

import (
	"fmt"
)

// START OMIT
type StartRequest struct {
	Image string
	Count *int
}

func Ptr[T any](v T) *T {
	return &v
}

func main() {
	req := StartRequest{
		Image: "debian:trixie",
		Count: Ptr(3), // HL
	}

	fmt.Println(req)
}

// END OMIT
