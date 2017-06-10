package main

import (
	"bytes"
	"net/http"

	"C"
)
import (
	"io"
)

//export fetch
func fetch(url *C.char) *C.char {
	resp, err := http.Get(C.GoString(url))
	if err != nil {
		return nil
	}

	var buf bytes.Buffer
	defer resp.Body.Close()
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil
	}
	return C.CString(buf.String())
}

func main() {}
