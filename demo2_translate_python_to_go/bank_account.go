package main

import (
	"errors"
	"fmt"
)

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-11
// BankAccount represents a simple checking account with deposit and withdraw functionality.
type BankAccount struct {
	name    string
	balance float64
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-11
// NewBankAccount creates a new BankAccount with the given name and initial balance.
// If no initial balance is provided, it defaults to 0.0.
func NewBankAccount(name string, balance ...float64) *BankAccount {
	initialBalance := 0.0
	if len(balance) > 0 {
		initialBalance = balance[0]
	}
	return &BankAccount{
		name:    name,
		balance: initialBalance,
	}
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-11
// Deposit adds the specified amount to the account balance.
// Returns an error if the amount is not positive.
func (ba *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit must be positive")
	}
	ba.balance += amount
	return nil
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-11
// Withdraw removes the specified amount from the account balance.
// Returns the withdrawn amount and an error if the amount is not positive or exceeds the balance.
func (ba *BankAccount) Withdraw(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("withdraw must be positive")
	}
	if amount > ba.balance {
		return 0, errors.New("insufficient funds")
	}
	ba.balance -= amount
	return amount, nil
}

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-11
// String returns a string representation of the account in the format "Name: $Balance".
// This implements the fmt.Stringer interface.
func (ba *BankAccount) String() string {
	return fmt.Sprintf("%s: $%.2f", ba.name, ba.balance)
}
