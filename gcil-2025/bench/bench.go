package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io"
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
// It converts the signature of slice.LoadLogs to loadFn in main.
func loadSlice(r io.Reader) iter.Seq[log.Log] {
	logs, err := slice.LoadLogs(r)
	if err != nil {
		return func(func(log.Log) bool) {
		}
	}

	return slices.Values(logs)
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
		fmt.Fprintf(os.Stderr, "Usage: %s slice|lazy\n", name)
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "error: missing method")
		os.Exit(1)
	}

	setupLogging()

	const fileName = "logs.json.gz"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	defer gz.Close()

	var loadFn func(r io.Reader) iter.Seq[log.Log]

	switch flag.Arg(0) {
	case "slice":
		loadFn = loadSlice
	case "lazy":
		loadFn = lazy.LoadLogs
	default:
		fmt.Fprintf(os.Stderr, "error: unknown loader - %q\n", flag.Arg(0))
		os.Exit(1)
	}

	start := time.Now()

	logs := loadFn(gz)

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
