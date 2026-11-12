package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type formatter interface {
	Row(columns []string, values []any) error
	Total(count int) error
	Error(err error, hint string)
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

func (jsonFormatter) Error(err error, hint string) {
	out := map[string]string{"error": err.Error()}
	if hint != "" {
		out["hint"] = hint
	}

	json.NewEncoder(os.Stderr).Encode(out)
}

func (textFormatter) Error(err error, hint string) {
	fmt.Fprintln(os.Stderr, "error:", err)
	if hint != "" {
		fmt.Fprintln(os.Stderr, "hint:", hint)
	}
}
