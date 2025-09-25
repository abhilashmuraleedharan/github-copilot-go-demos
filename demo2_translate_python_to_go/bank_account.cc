#include "bank_account.h"

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-11
BankAccount::BankAccount(const std::string& name, double balance)
    : name_(name), balance_(balance) {}

void BankAccount::Deposit(double amount) {
  if (amount <= 0) {
    throw std::invalid_argument("Deposit must be positive");
  }
  balance_ += amount;
}

double BankAccount::Withdraw(double amount) {
  if (amount <= 0) {
    throw std::invalid_argument("Withdraw must be positive");
  }
  if (amount > balance_) {
    throw std::invalid_argument("Insufficient funds");
  }
  balance_ -= amount;
  return amount;
}

std::string BankAccount::ToString() const {
  std::ostringstream oss;
  oss << name_ << ": $" << std::fixed << std::setprecision(2) << balance_;
  return oss.str();
}
