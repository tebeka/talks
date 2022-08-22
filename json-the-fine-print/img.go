package main

import (
	"encoding/json"
	"io/ioutil"
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
	data, _ := ioutil.ReadFile(name)

	img := Image{
		Name: name,
		Data: data,
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(img)
	// END_MAIN OMIT
}
