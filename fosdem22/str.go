package main

import (
	"encoding/json"
	"os"
)

func main() {
	// START_MAIN OMIT
	s := "<JSON>"
	json.NewEncoder(os.Stdout).Encode(s)
	// END_MAIN OMIT
}
