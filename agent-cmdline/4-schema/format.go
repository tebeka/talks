package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type formatter interface {
	Row(columns []string, values []any) error
	Total(count int) error
	Error(err error)
}

type textFormatter struct{}
type jsonFormatter struct{}

var format formatter = textFormatter{}

func (jsonFormatter) Row(columns []string, values []any) error {
	row := make(map[string]any, len(columns))
	for i, col := range columns {
		row[col] = values[i]
	}

	return json.NewEncoder(os.Stdout).Encode(row)
}

func (textFormatter) Row(columns []string, values []any) error {
	for i, col := range columns {
		val := display(values[i])
		fmt.Printf("%s: %v\n", col, val)
	}
	if len(columns) > 1 {
		fmt.Println()
	}

	return nil
}

func (jsonFormatter) Total(count int) error {
	return json.NewEncoder(os.Stdout).Encode(map[string]any{"_meta": map[string]int{"count": count}})
}

func (textFormatter) Total(count int) error {
	total := fmt.Sprintf("%d row(s)", count)
	fmt.Println(highlight(total))

	return nil
}

func (jsonFormatter) Error(err error) {
	out := map[string]string{"error": err.Error()}
	json.NewEncoder(os.Stderr).Encode(out)
}

func (textFormatter) Error(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
}
