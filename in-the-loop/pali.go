package main

import (
	"fmt"
)

func main() {
	for _, w := range []string{"d", "da", "dad"} {
		fmt.Println(w, isPalindrome(w))
	}
}

func isPalindrome(s string) bool {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}

	return true
}
