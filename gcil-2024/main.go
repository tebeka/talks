package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

//go:generate go run gen_db.go
//go:generate go fmt emoji_db.go

func matches(e Emoji, term string) bool {
	return strings.Contains(e.Name, term) || strings.Contains(e.Group, term)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s TERM\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	term := strings.ToLower(flag.Arg(0))
	found := false
	for _, e := range db {
		if !matches(e, term) {
			continue
		}
		found = true
		fmt.Printf("[%s] %s: %c\n", e.Group, e.Name, e.Char)
	}

	if !found {
		fmt.Fprintf(os.Stderr, "error: no match for '%q'\n", term)
		os.Exit(1)
	}
}
