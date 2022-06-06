package main

import (
	"fmt"
)

// START_TYPES OMIT
type syncer interface {
	Sync() error
}

type S3File struct{}

func (s S3File) Sync() error { // HL
	return nil
}

type File struct{}

func (f *File) Sync() error { // HL
	return nil
}

// END_TYPES OMIT

func main() {
	// START_MAIN OMIT
	var s syncer
	var s3 S3File

	s = &s3
	s = s3

	var file File
	s = &file
	s = file // HL
	// END_MAIN OMIT

	fmt.Println(s)
}
