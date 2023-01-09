package main

import (
	"fmt"
)

// START OMIT
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 { // HL
		if s[i] != s[j] {
			return false
		}
	}
	return true

}

// END OMIT

func main() {
	// START_M OMIT
	fmt.Println(isPalindrome("a"))
	fmt.Println(isPalindrome("ab"))
	fmt.Println(isPalindrome("aba"))
	// END_M OMIT
}
