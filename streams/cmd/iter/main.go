package main

import (
	"bufio"
	"errors"
	"fmt"
	"iter"
	"os"
	"path/filepath"
	"time"

	"logs/parser"
)

var ErrStop = errors.New("stop")

func iterLogs(root string) iter.Seq2[parser.Log, error] {
	fn := func(yield func(parser.Log, error) bool) {
		walkFn := func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				yield(parser.Log{}, err)
				return err
			}
			defer file.Close()

			s := bufio.NewScanner(file)
			for s.Scan() {
				log, err := parser.ParseLine(s.Text())
				if !yield(log, err) {
					return ErrStop
				}
			}
			if err := s.Err(); err != nil {
				yield(parser.Log{}, err)
				return err
			}

			return nil
		}

		filepath.Walk(root, walkFn)
	}

	return fn
}

func main() {
	var logs []parser.Log
	nErr := 0

	start := time.Now()
	for log, err := range iterLogs("logs") {
		if err != nil {
			nErr++
			continue
		}

		logs = append(logs, log)
	}

	duration := time.Since(start)
	fmt.Printf("%d logs (%d errors) in %v\n", len(logs), nErr, duration)
}
