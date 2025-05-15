package log

import "time"

type Log struct {
	Host    string    `json:"host"`
	Time    time.Time `json:"time"`
	Request string    `json:"request"`
	Status  int       `json:"status"`
	Bytes   int       `json:"bytes"`
}

func IsValid(log Log) bool {
	switch {
	case log.Host == "":
		return false
	case log.Time.IsZero():
		return false
	case log.Request == "":
		return false
	case log.Status < 100 || log.Status >= 600:
		return false
	case log.Bytes < 0:
		return false
	}

	return true
}
