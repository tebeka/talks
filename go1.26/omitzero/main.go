package main

import (
	"encoding/json"
	"os"
	"time"
)

// START OMIT
type LogQuery struct {
	Text      string    `json:"text,omitempty"`
	Level     string    `json:"level,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"` // HL
	EndTime   time.Time `json:"end_time,omitzero"`    // HL
}

func main() {
	query := LogQuery{
		Text:  "login",
		Level: "error",
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(query)
}

// END OMIT
