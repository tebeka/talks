package main

import (
	"flag"
	"fmt"
	"strconv"
)

// START_INTFLAG OMIT
// IntFlag is an integer flag with validation function
type IntFlag struct {
	check func(int) error
	ptr   *int
}

// Int return a new IntFlag
func NewIntFlag(ptr *int, check func(int) error) *IntFlag {
	return &IntFlag{ptr: ptr, check: check}
}

// END_INTFLAG OMIT

// START_VAR OMIT
func (v *IntFlag) String() string { // HL
	if v.ptr == nil {
		return ""
	}
	return fmt.Sprintf("%d", *v.ptr)
}

// Set value from a string
func (v *IntFlag) Set(s string) error { // HL
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	if v.check != nil {
		if err := v.check(i); err != nil { // HL
			return err
		}
	}

	*v.ptr = i
	return nil
}

// END_VAR OMIT

// START_MAIN OMIT
func checkRetries(n int) error {
	if n < 0 || n > 10 {
		return fmt.Errorf("retries = %d not in range [0:10]", n)
	}
	return nil
}

func main() {
	retries := 1

	flag.Var(NewIntFlag(&retries, checkRetries), "retries", "number of retries")
	flag.Parse()
	fmt.Println("retries:", retries)
}

// END_MAIN OMIT
