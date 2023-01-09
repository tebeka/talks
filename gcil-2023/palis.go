package main

import "fmt"

func main() {
	// START: OMIT
	vs := []string{"a", "ab", "aba"}
	type reply struct {
		s   string
		isP bool
	}

	// fan out
	ch := make(chan reply)
	for _, s := range vs {
		s := s // HL
		go func() {
			ch <- reply{s, isPalindrome(s)}
		}()
	}

	// collect
	for range vs { // HL
		r := <-ch
		fmt.Printf("%-5s -> %v\n", r.s, r.isP)
	}
	// END: OMIT
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 { // HL
		if s[i] != s[j] {
			return false
		}
	}
	return true

}
