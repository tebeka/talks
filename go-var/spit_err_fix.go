package main

import (
	"fmt"
	"log"

	// IMP_START OMIT
	"go.uber.org/multierr"
	// IMP_END OMIT
)

// Mock "os" for errors
type File struct{}

func (f File) Write([]byte) (int, error) {
	return 0, fmt.Errorf("NO WRITE FOR YOU!")
}

func (f File) Close() error {
	return nil
}

type OS struct{}

func (o OS) Create(string) (File, error) {
	return File{}, nil
}

var os OS

// SPIT_START OMIT
func spit(text string, fileName string) (err error) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("create: %s", err)
		return err
	}

	defer func() {
		err = multierr.Combine(err, file.Close()) // HL
	}()

	if _, err = file.Write([]byte(text)); err != nil {
		log.Printf("write: %s", err)
	}

	return
}

// SPIT_END OMIT

func main() {
	// MAIN_START OMIT
	if err := spit("Who's on first?", "sketch.txt"); err != nil {
		log.Fatal(err)
	}
	// MAIN_END OMIT
}
