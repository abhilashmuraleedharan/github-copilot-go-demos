// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-08-20
package bankaccount

import (
	"fmt"
	"errors"
)

// BankAccount represents a simple checking account with deposit and withdraw methods.
type BankAccount struct {
	name    string
	balance float64
}

// NewBankAccount creates a new BankAccount with the given name and optional initial balance.
func NewBankAccount(name string, balance float64) *BankAccount {
	return &BankAccount{name: name, balance: balance}
}

// Deposit adds the given amount to the account balance. Returns an error if amount is not positive.
func (a *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit must be positive")
	}
	a.balance += amount
	return nil
}

// Withdraw subtracts the given amount from the account balance. Returns an error if amount is not positive or insufficient funds.
func (a *BankAccount) Withdraw(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("withdraw must be positive")
	}
	if amount > a.balance {
		return 0, errors.New("insufficient funds")
	}
	a.balance -= amount
	return amount, nil
}

// String returns a string representation of the account.
func (a *BankAccount) String() string {
	return fmt.Sprintf("%s: $%.2f", a.name, a.balance)
}
