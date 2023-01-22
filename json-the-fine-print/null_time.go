package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	var t sql.NullTime

	data, _ := json.Marshal(time.Now())
	if err := json.Unmarshal(data, &t); err != nil {
		log.Fatal(err)
	}

	fmt.Println(t.Valid)
}
