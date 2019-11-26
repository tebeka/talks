package main

import (
	"encoding/json"
	"os"
)

type User struct {
	Login  string   `json:"login"`
	Name   string   `json:"name,omitempty"`
	Groups []string `json:"groups"`
	Token  string   `json:"-"`
}

func main() {
	u := User{
		Login:  "bugs",
		Groups: []string{"wheel", "network", "lp", "daffy"},
		Token:  "duck season",
	}
	json.NewEncoder(os.Stdout).Encode(u)
	// {"login":"bugs","groups":["wheel","network","lp","daffy"]}
}
