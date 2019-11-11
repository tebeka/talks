package main

import "fmt"

func init() { fmt.Printf("A ") }
func init() { fmt.Print("B ") }

func main() { fmt.Println() }
