package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// START_SLURP OMIT
func slurp(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}

// END_SLURP OMIT

func main() {
	// START_MAIN OMIT
	if data, err := slurp("song.txt"); err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Printf("%s\n", string(data))
	}
	// END_MAIN OMIT
}
