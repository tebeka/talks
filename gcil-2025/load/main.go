package main

import (
	"flag"
	"fmt"
	"iter"
	"log/slog"
	"os"
	"path"
	"runtime"
	"slices"
	"time"

	"mordor/lazy"
	"mordor/log"
	"mordor/slice"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// loadSlice returns a sequence of logs using simple.LoadLogs.
func loadSlice(root string) (iter.Seq[log.Log], error) {
	logs, err := slice.LoadLogs(root)
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
	flag.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage: %s simple|seq|mapper\n", name)
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "error: missing method")
		os.Exit(1)
	}

	var loadFn func(string) (iter.Seq[log.Log], error)

	switch flag.Arg(0) {
	case "slice":
		loadFn = loadSlice
	case "lazy":
		loadFn = lazy.LoadLogs
	default:
		fmt.Fprintf(os.Stderr, "error: unknown loader - %q\n", flag.Arg(0))
		os.Exit(1)
	}

	setupLogging()

	start := time.Now()

	logs, err := loadFn("logs.json.gz")
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
	alloc_mb := float64(mem.Alloc) / (1 << 20)

	p := message.NewPrinter(language.English)
	p.Printf("%d logs (%.2fsec, %.2fmb)\n", count, duration.Seconds(), alloc_mb)
}
