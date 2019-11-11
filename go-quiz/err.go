package main

import "fmt"

type OSError int

func (e *OSError) Error() string { return fmt.Sprintf("error #%d", e) }

func FileExists(path string) (bool, error) {
	var err *OSError
	return false, err // TODO
}

func main() {
	_, err := FileExists("/no/such/file")
	fmt.Println(err == nil)
}
