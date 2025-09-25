// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-17
#include "bank_account.h"
#include <gtest/gtest.h>
#include <stdexcept>

class BankAccountTest : public ::testing::Test {
 protected:
  void SetUp() override {
    // Set up common test data
  }

  void TearDown() override {
    // Clean up after each test
  }
};

// Test constructor with default balance
TEST_F(BankAccountTest, ConstructorWithDefaultBalance) {
  BankAccount account("John Doe");
  EXPECT_EQ(account.ToString(), "John Doe: $0.00");
}

// Test constructor with initial balance
TEST_F(BankAccountTest, ConstructorWithInitialBalance) {
  BankAccount account("Jane Smith", 1000.50);
  EXPECT_EQ(account.ToString(), "Jane Smith: $1000.50");
}

// Test constructor with zero balance explicitly
TEST_F(BankAccountTest, ConstructorWithZeroBalance) {
  BankAccount account("Zero Balance", 0.0);
  EXPECT_EQ(account.ToString(), "Zero Balance: $0.00");
}

// Test successful deposit
TEST_F(BankAccountTest, DepositValidAmount) {
  BankAccount account("Test User", 100.0);
  account.Deposit(50.25);
  EXPECT_EQ(account.ToString(), "Test User: $150.25");
}

// Test multiple deposits
TEST_F(BankAccountTest, MultipleDeposits) {
  BankAccount account("Multi Deposit", 0.0);
  account.Deposit(100.0);
  account.Deposit(50.50);
  account.Deposit(25.75);
  EXPECT_EQ(account.ToString(), "Multi Deposit: $176.25");
}

// Test deposit with zero amount (should throw exception)
TEST_F(BankAccountTest, DepositZeroAmount) {
  BankAccount account("Test User", 100.0);
  EXPECT_THROW(account.Deposit(0.0), std::invalid_argument);
}

// Test deposit with negative amount (should throw exception)
TEST_F(BankAccountTest, DepositNegativeAmount) {
  BankAccount account("Test User", 100.0);
  EXPECT_THROW(account.Deposit(-50.0), std::invalid_argument);
}

// Test deposit exception message
TEST_F(BankAccountTest, DepositExceptionMessage) {
  BankAccount account("Test User", 100.0);
  try {
    account.Deposit(-10.0);
    FAIL() << "Expected std::invalid_argument";
  } catch (const std::invalid_argument& e) {
    EXPECT_STREQ(e.what(), "Deposit must be positive");
  }
}

// Test successful withdrawal
TEST_F(BankAccountTest, WithdrawValidAmount) {
  BankAccount account("Test User", 100.0);
  double withdrawn = account.Withdraw(30.50);
  EXPECT_DOUBLE_EQ(withdrawn, 30.50);
  EXPECT_EQ(account.ToString(), "Test User: $69.50");
}

// Test withdrawal of entire balance
TEST_F(BankAccountTest, WithdrawEntireBalance) {
  BankAccount account("Test User", 500.75);
  double withdrawn = account.Withdraw(500.75);
  EXPECT_DOUBLE_EQ(withdrawn, 500.75);
  EXPECT_EQ(account.ToString(), "Test User: $0.00");
}

// Test multiple withdrawals
TEST_F(BankAccountTest, MultipleWithdrawals) {
  BankAccount account("Multi Withdraw", 1000.0);
  account.Withdraw(200.0);
  account.Withdraw(150.50);
  account.Withdraw(100.25);
  EXPECT_EQ(account.ToString(), "Multi Withdraw: $549.25");
}

// Test withdrawal with zero amount (should throw exception)
TEST_F(BankAccountTest, WithdrawZeroAmount) {
  BankAccount account("Test User", 100.0);
  EXPECT_THROW(account.Withdraw(0.0), std::invalid_argument);
}

// Test withdrawal with negative amount (should throw exception)
TEST_F(BankAccountTest, WithdrawNegativeAmount) {
  BankAccount account("Test User", 100.0);
  EXPECT_THROW(account.Withdraw(-50.0), std::invalid_argument);
}

// Test withdrawal exceeding balance (insufficient funds)
TEST_F(BankAccountTest, WithdrawInsufficientFunds) {
  BankAccount account("Test User", 50.0);
  EXPECT_THROW(account.Withdraw(100.0), std::invalid_argument);
}

// Test withdrawal exception messages
TEST_F(BankAccountTest, WithdrawNegativeExceptionMessage) {
  BankAccount account("Test User", 100.0);
  try {
    account.Withdraw(-10.0);
    FAIL() << "Expected std::invalid_argument";
  } catch (const std::invalid_argument& e) {
    EXPECT_STREQ(e.what(), "Withdraw must be positive");
  }
}

TEST_F(BankAccountTest, WithdrawInsufficientFundsMessage) {
  BankAccount account("Test User", 50.0);
  try {
    account.Withdraw(100.0);
    FAIL() << "Expected std::invalid_argument";
  } catch (const std::invalid_argument& e) {
    EXPECT_STREQ(e.what(), "Insufficient funds");
  }
}

// Test ToString method formatting
TEST_F(BankAccountTest, ToStringFormatting) {
  BankAccount account("Format Test", 123.456);
  EXPECT_EQ(account.ToString(), "Format Test: $123.46");
}

// Test ToString with large balance
TEST_F(BankAccountTest, ToStringLargeBalance) {
  BankAccount account("Rich User", 1234567.89);
  EXPECT_EQ(account.ToString(), "Rich User: $1234567.89");
}

// Test ToString with fractional cents (rounding)
TEST_F(BankAccountTest, ToStringRounding) {
  BankAccount account("Rounding Test", 99.999);
  EXPECT_EQ(account.ToString(), "Rounding Test: $100.00");
}

// Test combined operations
TEST_F(BankAccountTest, CombinedOperations) {
  BankAccount account("Combined Test", 1000.0);
  
  // Deposit some money
  account.Deposit(250.75);
  EXPECT_EQ(account.ToString(), "Combined Test: $1250.75");
  
  // Withdraw some money
  double withdrawn = account.Withdraw(500.25);
  EXPECT_DOUBLE_EQ(withdrawn, 500.25);
  EXPECT_EQ(account.ToString(), "Combined Test: $750.50");
  
  // Another deposit
  account.Deposit(100.00);
  EXPECT_EQ(account.ToString(), "Combined Test: $850.50");
}

// Test with special characters in name
TEST_F(BankAccountTest, SpecialCharactersInName) {
  BankAccount account("John O'Connor-Smith", 100.0);
  EXPECT_EQ(account.ToString(), "John O'Connor-Smith: $100.00");
}

// Test with empty name
TEST_F(BankAccountTest, EmptyName) {
  BankAccount account("", 100.0);
  EXPECT_EQ(account.ToString(), ": $100.00");
}

// Test precision with very small amounts
TEST_F(BankAccountTest, SmallAmountPrecision) {
  BankAccount account("Precision Test", 0.01);
  EXPECT_EQ(account.ToString(), "Precision Test: $0.01");
  
  account.Deposit(0.01);
  EXPECT_EQ(account.ToString(), "Precision Test: $0.02");
}