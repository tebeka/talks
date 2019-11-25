package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	i := complex(0, 1)
	_, err := json.Marshal(i)
	fmt.Println(err) // json: unsupported type: complex128

	data, _ := json.Marshal(time.Now())
	var out interface{}
	json.Unmarshal(data, &out)
	fmt.Printf("%T\n", out) // string

	n := 20
	data, _ = json.Marshal(n)
	json.Unmarshal(data, &out)
	fmt.Printf("%T\n", out) // float64
}
