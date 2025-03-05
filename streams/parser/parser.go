package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// 199.72.81.55 - - [01/Jul/1995:00:00:01 -0400] "GET /history/apollo/ HTTP/1.0" 200 6245
// unicomp6.unicomp.net - - [01/Jul/1995:00:00:06 -0400] "GET /shuttle/countdown/ HTTP/1.0" 200 3985
// logs/2/x06:91503:70:klothos.crl.research.digital.com - - [10/Jul/1995:16:45:50 -0400] "" 400 -
var logRegexp = regexp.MustCompile(`^(\S+) - - \[(\d{2}/\w{3}/\d{4}:\d{2}:\d{2}:\d{2}) -0400\] "([^"]*)" (\d{3}) (\S+)$`)

type Log struct {
	Host    string
	Time    time.Time
	Request string
	Status  int
	Bytes   int
}

func ParseLine(line string) (Log, error) {
	matches := logRegexp.FindStringSubmatch(line)
	if matches == nil {
		return Log{}, fmt.Errorf("Could not parse line: %s\n", line)
	}

	log := Log{
		Host:    matches[1],
		Request: matches[3],
	}

	// We know conversion is safe because of the regexp
	log.Time, _ = time.Parse("02/Jan/2006:15:04:05", matches[2])
	log.Status, _ = strconv.Atoi(matches[4])
	if matches[5] != "-" {
		log.Bytes, _ = strconv.Atoi(matches[5])
	}

	return log, nil
}
