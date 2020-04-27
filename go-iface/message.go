package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

// Message format: [type:3 bytes][payload size: 8 bytes, uint64][payload]

var (
	typeLen   = 3
	sizeLen   = binary.Size(uint64(0))
	headerLen = typeLen + sizeLen
	encoder   = binary.BigEndian
)

type Message interface {
	Type() string
}

// Marshal to byte array
func Marshal(m Message) ([]byte, error) {
	typ := m.Type()
	if len(typ) != typeLen {
		return nil, fmt.Errorf("bad type: %q", typ)
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(m); err != nil {
		return nil, err
	}
	data := buf.Bytes()

	msg := make([]byte, headerLen+len(data))
	copy(msg, []byte(typ))
	size := uint64(len(data))
	encoder.PutUint64(msg[typeLen:], size)
	copy(msg[headerLen:], data)

	return msg, nil
}

// Metric is a metric message
type Metric struct {
	Time  time.Time
	Name  string
	Value float64
}

func (m Metric) Type() string {
	return "MET"
}

// Log is a log message
type Log struct {
	Time    time.Time
	Level   int
	Message string
}

func (l Log) Type() string {
	return "LOG"
}

func main() {
	m := Metric{time.Now(), "Memory", 329_093}
	data, err := Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	l := Log{time.Now(), 7, "The world is round"}
	data, err = Marshal(l)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
