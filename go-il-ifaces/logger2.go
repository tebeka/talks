package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Level byte

const (
	Debug Level = iota + 1
	Info
	Error
)

func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Error:
		return "ERROR"
	}

	return fmt.Sprintf("Level <%d>", l)
}

// START_SYNCER OMIT
type syncer interface {
	Sync() error
}

type Logger struct {
	level Level
	w     io.Writer
	s     syncer // HL
}

func NewLogger(level Level, out io.Writer) Logger {
	log := Logger{level, out, nil}
	if s, ok := out.(syncer); ok { // HL
		log.s = s // HL
	} // HL
	return log
}

// END_SYNCER OMIT

// START_LOG OMIT
func (l Logger) log(level Level, format string, args ...any) {
	if l.level > level {
		return
	}

	msg := fmt.Sprintf(format, args...)
	ts := time.Now().UTC().Format(time.RFC3339)
	fmt.Fprintf(l.w, "[%s] - %s - %s\n", ts, level, msg)

	if l.s != nil { // HL
		l.s.Sync() // HL
	} // HL
}

// END_LOG OMIT

func (l Logger) Debug(format string, args ...any) {
	l.log(Debug, format, args...)
}

func (l Logger) Info(format string, args ...any) {
	l.log(Info, format, args...)
}

func (l Logger) Error(format string, args ...any) {
	l.log(Error, format, args...)
}

func main() {
	log := NewLogger(Info, os.Stdout)
	log.Debug("debug")
	log.Info("info")
	log.Error("error")
}
