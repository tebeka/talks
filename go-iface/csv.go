package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var (
	csvHeader = []string{"time", "name", "value"}
)

// Metric is a metric
type Metric struct {
	Time  time.Time
	Name  string
	Value float64
}

// Writer writes CSV records
type Writer struct {
	wtr         io.Writer
	csv         *csv.Writer
	wroteHeader bool
	row         []string
}

// NewWriter returns a new CSV Writer
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		wtr: w,
		csv: csv.NewWriter(w),
		row: make([]string, len(csvHeader)),
	}

}

// Append appends metrics to the CSV
func (w *Writer) Append(metrics []Metric) error {
	if !w.wroteHeader {
		if err := w.csv.Write(csvHeader); err != nil {
			return err
		}
		w.wroteHeader = true
	}

	for _, m := range metrics {
		w.row[0] = m.Time.Format(time.RFC3339)
		w.row[1] = m.Name
		w.row[2] = fmt.Sprintf("%f", m.Value)
		if err := w.csv.Write(w.row); err != nil {
			return err
		}
	}
	return nil
}

// File Sync
type syncer interface {
	Sync() error
}

// Wait waits for write to be finished
func (w *Writer) Wait() error {
	w.csv.Flush()
	if err := w.csv.Error(); err != nil {
		return err
	}

	if s, ok := w.wtr.(syncer); ok {
		return s.Sync()
	}

	return nil
}

func main() {
	const count = 10
	start := time.Now().Add(-count * time.Second)
	var metrics []Metric
	for i := 0; i < count; i++ {
		m := Metric{
			Time:  start.Add(time.Duration(i) * time.Second),
			Name:  "CPU",
			Value: rand.Float64() * 100,
		}
		metrics = append(metrics, m)
	}

	w := NewWriter(os.Stdout)
	w.Append(metrics)
	w.Wait()
}
