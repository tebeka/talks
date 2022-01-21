package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// START_IMAGE OMIT
type Image struct {
	Name string
	Data []byte // HL
}

// END_IMAGE OMIT

func main() {
	// START_MAIN OMIT
	name := "4.png"
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	img := Image{
		Name: name,
		Data: data,
	}

	if err := json.NewEncoder(os.Stdout).Encode(img); err != nil {
		log.Fatal(err)
	}
	// END_MAIN OMIT
}
