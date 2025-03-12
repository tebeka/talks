package main

import (
	"flag"
	"fmt"
	"iter"
	"log/slog"
	"os"
	"runtime"
	"slices"
	"time"

	"logs/loader/mapper"
	"logs/loader/seq"
	"logs/loader/simple"
	"logs/parser"
)

// loadSimple returns a sequence of logs using simple.LoadLogs.
func loadSimple(root string) (iter.Seq[parser.Log], error) {
	logs, err := simple.LoadLogs(root)
	if err != nil {
		return nil, err
	}

	return slices.Values(logs), nil
}

func setupLogging() {
	h := slog.NewTextHandler(
		os.Stderr,
		&slog.HandlerOptions{Level: slog.LevelError},
	)
	logger := slog.New(h)
	slog.SetDefault(logger)

}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "error: missing method")
		os.Exit(1)
	}

	var loadFn func(string) (iter.Seq[parser.Log], error)

	switch flag.Arg(0) {
	case "simple":
		loadFn = loadSimple
	case "seq":
		loadFn = seq.LoadLogs
	case "mapper":
		loadFn = mapper.LoadLogs
	default:
		fmt.Fprintf(os.Stderr, "error: unknown loader - %q\n", flag.Arg(0))
		os.Exit(1)
	}

	setupLogging()

	start := time.Now()

	logs, err := loadFn("logs")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	count := 0
	for range logs {
		count++
	}

	duration := time.Since(start)

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	alloc_mb := mem.Alloc / (1 << 20)

	fmt.Printf("%d logs (%v %dmb)\n", count, duration, alloc_mb)
}
