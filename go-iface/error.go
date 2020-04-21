package main

import (
	"fmt"
)

type OSError int

func (e *OSError) Error() string {
	return fmt.Sprintf("error #%d", *e)
}

func FileExists(path string) (bool, error) {
	var err *OSError
	return false, err // TODO
}

func main() {
	if _, err := FileExists("/no/such/file"); err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Println("OK")
	}
}
