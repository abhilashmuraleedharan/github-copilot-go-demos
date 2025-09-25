# BankAccount Unit Tests

This directory contains unit tests for the `BankAccount` C++ class using Google Test framework.

## Files

- `bank_account.h` - Header file for the BankAccount class
- `bank_account.cc` - Implementation of the BankAccount class
- `bank_account_test.cc` - Unit tests for the BankAccount class
- `CMakeLists.txt` - CMake build configuration
- `Makefile` - Alternative build configuration using Make

## Test Coverage

The unit tests cover the following scenarios:

### Constructor Tests
- Default balance (0.0)
- Initial balance provided
- Zero balance explicitly set
- Special characters in account name
- Empty account name

### Deposit Tests
- Valid positive deposits
- Multiple consecutive deposits
- Zero amount deposit (exception expected)
- Negative amount deposit (exception expected)
- Exception message validation

### Withdrawal Tests
- Valid withdrawals within balance
- Withdrawal of entire balance
- Multiple consecutive withdrawals
- Zero amount withdrawal (exception expected)
- Negative amount withdrawal (exception expected)
- Insufficient funds withdrawal (exception expected)
- Exception message validation

### ToString Tests
- Proper formatting with 2 decimal places
- Large balance formatting
- Rounding behavior
- Special characters in names

### Integration Tests
- Combined deposit and withdrawal operations
- Precision with small amounts

## Building and Running Tests

### Option 1: Using CMake

```bash
# Create build directory
mkdir build
cd build

# Configure and build
cmake ..
make

# Run tests
./bank_account_test
```

### Option 2: Using Make

```bash
# Build the tests
make

# Run the tests
make test

# Clean build artifacts
make clean
```

### Option 3: Manual Compilation

```bash
# Compile and link (assuming Google Test is installed)
g++ -std=c++17 -Wall -Wextra -g -o bank_account_test bank_account.cc bank_account_test.cc -lgtest -lgtest_main -pthread

# Run tests
./bank_account_test
```

## Prerequisites

- C++17 compatible compiler (g++, clang++)
- Google Test framework
- CMake 3.14+ (if using CMake)

### Installing Google Test

#### Ubuntu/Debian:
```bash
sudo apt-get update
sudo apt-get install libgtest-dev cmake
cd /usr/src/gtest
sudo cmake .
sudo make
sudo cp *.a /usr/lib
```

#### macOS (using Homebrew):
```bash
brew install googletest
```

#### Windows:
- Download and build Google Test from source, or
- Use vcpkg: `vcpkg install gtest`

## Test Output

When run successfully, you should see output similar to:
```
[==========] Running 23 tests from 1 test suite.
[----------] Global test environment set-up.
[----------] 23 tests from BankAccountTest
[ RUN      ] BankAccountTest.ConstructorWithDefaultBalance
[       OK ] BankAccountTest.ConstructorWithDefaultBalance (0 ms)
...
[----------] 23 tests from BankAccountTest (X ms total)

[----------] Global test environment tear-down
[==========] 23 tests from 1 test suite ran. (X ms total)
[  PASSED  ] 23 tests.
```