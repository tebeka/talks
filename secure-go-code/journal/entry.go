package main

import "time"

type Entry struct {
	Time    time.Time `json:"time"`
	User    string    `json:"user"`
	Content string    `json:"content"`
}
