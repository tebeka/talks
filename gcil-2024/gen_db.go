//go:build ignore

// Data source https://unicode.org/Public/emoji/15.1/emoji-sequences.txt

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var groups = map[string]string{
	"Basic_Emoji":                 "basic",
	"Emoji_Keycap_Sequence":       "keycap",
	"RGI_Emoji_Modifier_Sequence": "modifier",
	"RGI_Emoji_Tag_Sequence":      "tag",
	"RGI_Emoji_Flag_Sequence":     "flag",
}

// 231A..231B    ; Basic_Emoji                  ; watch..hourglass done                                          # E0.6   [2] (⌚..⌛)
// 23F0          ; Basic_Emoji                  ; alarm clock                                                    # E0.6   [1] (⏰)
// FIXME: Doesn't match
// 0031 FE0F 20E3; Emoji_Keycap_Sequence        ; keycap: 1                                                      # E0.6   [1] (1️⃣)
var lineRe = regexp.MustCompile(`.*; ([^;]+);\s*([a-zA-Z :]+)(\.\.[a-zA-Z ]+)?.*\((.)(\.\..)?`)

func parseLine(line string) (string, []string, []rune, error) {
	matches := lineRe.FindStringSubmatch(line)
	if matches == nil {
		return "", nil, nil, nil
	}

	names := [2]string{}
	runes := [2]rune{}
	count := 1
	group := groups[strings.TrimSpace(matches[1])]
	if group == "" {
		return "", nil, nil, fmt.Errorf("unknown group %q", matches[1])
	}
	names[0] = strings.ToLower(strings.TrimSpace(matches[2]))
	if matches[3] != "" {
		names[1] = strings.ToLower(strings.TrimSpace(matches[3][2:])) // trim ..
		count = 2
	}
	runes[0] = []rune(matches[4])[0]
	if matches[5] != "" {
		runes[1] = []rune(matches[5])[2] // ignore ..
	}

	return group, names[:count], runes[:count], nil
}

var header = `package main

type Emoji struct {
	Name  string
	Group string
	Char  rune
}

var db = []Emoji{
`

func main() {
	file, err := os.Open("emoji-sequences.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	defer file.Close()

	outFile := "emoji_db.go"
	out, err := os.Create(outFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	defer out.Close()

	fmt.Fprint(out, header)

	lnum := 0
	s := bufio.NewScanner(file)
	n := 0
	for s.Scan() {
		lnum++
		group, names, runes, err := parseLine(s.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %d: %v\n", lnum, err)
			os.Exit(1)
		}
		if group == "" { // not an info line
			continue
		}

		for i := range names {
			fmt.Fprintf(out, "\t{ %q, %q, '%c' },\n", names[i], group, runes[i])
		}
		n += len(names)
	}
	if err := s.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	fmt.Fprintln(out, "}")

	fmt.Printf("%s: %d emojies\n", outFile, n)
}
