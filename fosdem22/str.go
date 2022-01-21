package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	// START_MAIN OMIT
	s := "<JSON>"
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(s); err != nil {
		log.Fatal(err)
	}
	// END_MAIN OMIT
}
