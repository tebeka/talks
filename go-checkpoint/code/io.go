package io

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

func Copy(dst Writer, src Reader) (written int64, err error)
