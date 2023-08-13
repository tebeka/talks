package main

import (
    "fmt"
)

const (
    Read = 1 << iota
    Write
    Execute
)

func main() {
    fmt.Println(Execute)
}
