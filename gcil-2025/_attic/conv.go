package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"logs/loader/mapper"
	"os"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func setupLogging() {
	h := slog.NewTextHandler(
		os.Stderr,
		&slog.HandlerOptions{Level: slog.LevelError},
	)
	logger := slog.New(h)
	slog.SetDefault(logger)

}

func main() {
	setupLogging()

	logs, err := mapper.LoadLogs("logs")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	out, err := os.Create("logs.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer out.Close()

	enc := json.NewEncoder(out)
	count := 0
	for log := range logs {
		count++
		if err := enc.Encode(log); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	}

	p := message.NewPrinter(language.English)
	p.Printf("%d logs\n", count)
}
