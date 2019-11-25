package main

import (
	"encoding/json"
	"fmt"
)

type Job struct {
	User   string
	Action string
	Count  int
}

func main() {
	data := []byte(`{
		"user": "saitama",
		"action": "punch"
	}`)

	var job Job
	json.Unmarshal(data, &job)
	fmt.Printf("%+v\n", job) // {User:saitama Action:punch Count:0}
}
