package main

import (
	"encoding/json"
	"os"
)

func main() {
	// START_MAIN OMIT
	s := "<JSON>"
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(s)
	// END_MAIN OMIT
}