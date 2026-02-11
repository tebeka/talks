package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// START OMIT
	file, err := os.Open("/var/log/app.log")
	if err != nil {
		if fsErr, ok := errors.AsType[*fs.PathError](err); ok { // HL
			fmt.Printf("error: path error (%v)\n", fsErr)
		}
	}
	defer file.Close()
	// END OMIT
}
