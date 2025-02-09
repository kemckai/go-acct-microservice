package main

import (
	"fmt"
)

type Account struct {
	ID      int
	Name    string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
	fmt.Printf("Deposited %.2f to account %d. New balance: %.2f\n", amount, a.ID, a.Balance)
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.Balance {
		return fmt.Errorf("insufficient funds")
	}
	a.Balance -= amount
	fmt.Printf("Withdrew %.2f from account %d. New balance: %.2f\n", amount, a.ID, a.Balance)
	return nil
}

func main() {
	account := Account{ID: 1, Name: "John Doe", Balance: 1000.0}
	account.Deposit(500.0)
	err := account.Withdraw(200.0)
	if err != nil {
		fmt.Println(err)
	}
}