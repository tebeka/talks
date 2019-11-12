package main

import (
	"bufio"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	// *B:* panic
	bulletRe = regexp.MustCompile(`\*[ABCD]:\* `)
)

type TestCase struct {
	file    string
	answers []string
}

func parseSlides(path string, t *testing.T) []*TestCase {
	require := require.New(t)
	file, err := os.Open(path)
	require.NoError(err, "open slides")

	s := bufio.NewScanner(file)
	var cases []*TestCase
	var tc *TestCase
	lnum := 0
	for s.Scan() {
		lnum++
		line := s.Text()
		if len(line) > 0 && line[0] == '#' {
			continue
		}
		// .play -edit -numbers empty-map-ok.go
		if strings.Contains(s.Text(), ".play") {
			if tc != nil {
				require.Equalf(4, len(tc.answers), "case before %d", lnum)
				cases = append(cases, tc)
			}
			fields := strings.Fields(line)
			tc = &TestCase{file: fields[len(fields)-1]}
			continue
		}

		match := bulletRe.FindString(line)
		if match == "" {
			continue
		}
		text := line[len(match):]
		require.Truef(tc != nil, "%d: bullet not in section", lnum)
		tc.answers = append(tc.answers, text)
	}
	require.NoError(s.Err(), "scan")

	if tc != nil {
		require.Equalf(4, len(tc.answers), "last case")
		cases = append(cases, tc)
	}
	return cases
}

func hasAnswer(output string, answers []string) bool {
	for _, ans := range answers {
		if strings.Contains(output, ans) {
			return true
		}
	}
	return false
}

func hasFail(answers []string) bool {
	for _, ans := range answers {
		for _, text := range []string{"panic", "Won't compile", "Deadlock"} {
			if strings.Contains(ans, text) {
				return true
			}
		}
	}
	return true
}

func TestSlides(t *testing.T) {
	cases := parseSlides("quiz.slide", t)
	for _, tc := range cases {
		t.Run(tc.file, func(t *testing.T) {
			require := require.New(t)
			cmd := exec.Command("go", "run", tc.file)
			data, err := cmd.CombinedOutput()
			output := string(data)
			if err == nil {
				require.True(hasAnswer(output, tc.answers))
			} else {
				require.True(hasFail(tc.answers), "fail")
			}
		})
	}
}
