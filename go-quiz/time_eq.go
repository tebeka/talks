package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	data, _ := json.Marshal(t1)

	var t2 time.Time
	_ = json.Unmarshal(data, &t2)

	fmt.Println(t1 == t2)
}
