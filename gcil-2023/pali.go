package main

import (
	"fmt"
)

// START OMIT
func isPalindrome(s string) bool {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 { // HL
		if chars[i] != chars[j] {
			return false
		}
	}

	return true
}

// END OMIT

func main() {
	// START_M OMIT
	fmt.Println(isPalindrome("a"))
	fmt.Println(isPalindrome("a♡"))
	fmt.Println(isPalindrome("a♡a"))
	// END_M OMIT
}
