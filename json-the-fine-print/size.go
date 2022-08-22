package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

func main() {
	// START_MAIN OMIT
	var b byte = 106 // j
	data, _ := json.Marshal(b)

	fmt.Println("Go  :", unsafe.Sizeof(b))
	fmt.Println("JSON:", len(data))
	// END_MAIN OMIT
}
