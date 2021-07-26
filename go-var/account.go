package main

import "fmt"

/*
padding

*/

// START OMIT
func makeAccount(balance int) func(int) int { // HL
	fn := func(amount int) int {
		balance += amount // HL
		return balance
	}

	return fn
}

// END OMIT

func main() {
	// MAIN_START OMIT
	acct := makeAccount(100)
	fmt.Println(acct(100))
	fmt.Println(acct(-50))
	// MAIN_END OMIT
}
