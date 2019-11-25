package main

import (
	"encoding/json"
	"fmt"
)

type Job struct {
	User   string
	Action string
	JCount *int `json:"count"`
}

func (j *Job) Count() int {
	if j.JCount == nil {
		return 1
	}
	return *j.JCount
}

func main() {
	data := []byte(`{
		"user": "saitama",
		"action": "punch"
	}`)

	var job Job
	json.Unmarshal(data, &job)
	fmt.Println(job.Count())
}
