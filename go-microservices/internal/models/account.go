package models

import "fmt"

type Account struct {
	ID      int
	Name    string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.Balance {
		return fmt.Errorf("insufficient funds")
		
	}
	a.Balance -= amount
	return nil
}