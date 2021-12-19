package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var msg interface{} = 1
	data, _ := json.Marshal(msg)
	var reply interface{}
	json.Unmarshal(data, &reply)
	fmt.Println(msg, reply, msg == reply)
}
