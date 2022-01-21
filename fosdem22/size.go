package main

import (
	"encoding/json"
	"fmt"
	"log"
	"unsafe"
)

func main() {
	// START_MAIN OMIT
	var n int8 = 123
	data, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Go  :", unsafe.Sizeof(n))
	fmt.Println("JSON:", len(data))
	// END_MAIN OMIT
}
