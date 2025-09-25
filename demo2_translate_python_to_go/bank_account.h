#ifndef DEMO2_TRANSLATE_PYTHON_TO_GO_BANK_ACCOUNT_H_
#define DEMO2_TRANSLATE_PYTHON_TO_GO_BANK_ACCOUNT_H_

#include <string>
#include <stdexcept>
#include <sstream>
#include <iomanip>

// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-01-11
// BankAccount models a simple checking account with deposit and withdraw operations.
class BankAccount {
 public:
  // Constructs a BankAccount with the given name and optional initial balance.
  explicit BankAccount(const std::string& name, double balance = 0.0);

  // Deposits a positive amount into the account.
  // Throws std::invalid_argument if amount is not positive.
  void Deposit(double amount);

  // Withdraws a positive amount from the account.
  // Throws std::invalid_argument if amount is not positive or insufficient funds.
  double Withdraw(double amount);

  // Returns a string representation of the account in the format "Name: $Balance".
  std::string ToString() const;

 private:
  std::string name_;
  double balance_;
};

#endif  // DEMO2_TRANSLATE_PYTHON_TO_GO_BANK_ACCOUNT_H_
