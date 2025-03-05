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

func iterLines(root string) iter.Seq2[string, error] {
	fn := func(yield func(string, error) bool) {
		walkFn := func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				yield("", err)
				return err
			}
			defer file.Close()

			s := bufio.NewScanner(file)
			for s.Scan() {
				if !yield(s.Text(), err) {
					return ErrStop
				}
			}
			if err := s.Err(); err != nil {
				yield("", err)
				return err
			}

			return nil
		}

		filepath.Walk(root, walkFn)
	}

	return fn
}

func Map[F any, T any](seq iter.Seq2[F, error], mapper func(F) (T, error)) iter.Seq2[T, error] {
	fn := func(yield func(T, error) bool) {
		for v, err := range seq {
			var t T
			if err == nil {
				t, err = mapper(v)
			}
			if !yield(t, err) {
				return
			}
		}
	}

	return fn
}

func iterLogs(root string) iter.Seq2[parser.Log, error] {
	return Map(iterLines(root), parser.ParseLine)
}

func Filter[T any](seq iter.Seq2[T, error], pred func(T) bool) iter.Seq2[T, error] {
	fn := func(yield func(T, error) bool) {
		for v, err := range seq {
			if err != nil || pred(v) {
				if !yield(v, err) {
					return
				}
			}
		}
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
