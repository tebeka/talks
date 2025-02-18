package main

import "fmt"

func main() {
	type Account struct {
		Login   string
		Type    string
		Balance int // ¢
	}

	bank := []Account{
		{"donald", "regular", 100},
		{"scrooge", "vip", 1_000_000},
	}

	// Give VIP accounts 1000 bonus
	for _, a := range bank {
		if a.Type == "vip" {
			a.Balance += 1_000
		}
	}

	fmt.Println(bank)
}
