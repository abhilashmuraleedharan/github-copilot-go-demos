package bank

import (
    "errors"
    "fmt"
)

type BankAccount struct {
    name    string
    balance float64
}

func NewBankAccount(name string, initialBalance float64) *BankAccount {
    return &BankAccount{name: name, balance: initialBalance}
}

func (b *BankAccount) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("Deposit must be positive")
    }
    b.balance += amount
    return nil
}

func (b *BankAccount) Withdraw(amount float64) (float64, error) {
    if amount <= 0 {
        return 0, errors.New("Withdraw must be positive")
    }
    if amount > b.balance {
        return 0, errors.New("Insufficient funds")
    }
    b.balance -= amount
    return amount, nil
}

func (b *BankAccount) String() string {
    return fmt.Sprintf("%s: $%.2f", b.name, b.balance)
}
