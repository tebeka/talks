package main

import (
	"fmt"
	"regexp"
)

// Apache/NCSA Combined Log Format Regex
var logRegex = regexp.MustCompile(`^(\S+) (\S+) (\S+) \[([^\]]+)\] "([^"]*)" (\d+) (\d+)$`)

type LogEntry struct {
	Host       string
	Identity   string
	User       string
	Timestamp  string
	Request    string
	StatusCode int
	BytesSent  int
}

func parseLogLine(line string) (*LogEntry, error) {
	matches := logRegex.FindStringSubmatch(line)
	if matches == nil {
		return nil, fmt.Errorf("could not parse log line: %s", line)
	}

	// Convert status code and bytes sent to integers
	statusCode, _ := strconv.Atoi(matches[6])
	bytesSent, _ := strconv.Atoi(matches[7])

	return &LogEntry{
		Host:       matches[1],
		Identity:   matches[2],
		User:       matches[3],
		Timestamp:  matches[4],
		Request:    matches[5],
		StatusCode: statusCode,
		BytesSent:  bytesSent,
	}, nil
}

func main() {
	// Example log lines
	logLines := []string{
		`199.72.81.55 - - [01/Jul/1995:00:00:01 -0400] "GET /history/apollo/ HTTP/1.0" 200 6245`,
		`unicomp6.unicomp.net - - [01/Jul/1995:00:00:06 -0400] "GET /shuttle/countdown/ HTTP/1.0" 200 3985`,
	}

	for _, line := range logLines {
		entry, err := parseLogLine(line)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Host: %s, Request: %s, Status: %d, Bytes: %d\n", 
			entry.Host, entry.Request, entry.StatusCode, entry.BytesSent)
	}
}
