package main

import "fmt"

type Account struct {
	Owner   string
	Balance float64
}

func (a *Account) Deposit(sum float64) {
	a.Balance += sum
}

func main() {
	a := Account{"Scrooge McDuck", 1000}
	Account.Deposit(&a, 12.3)
	fmt.Println(a)
}
