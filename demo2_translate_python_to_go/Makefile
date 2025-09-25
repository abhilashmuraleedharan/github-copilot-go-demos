# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-17
# Makefile for BankAccount unit tests

CXX = g++
CXXFLAGS = -std=c++17 -Wall -Wextra -g
GTEST_FLAGS = -lgtest -lgtest_main -pthread

# Source files
SOURCES = bank_account.cc bank_account_test.cc
TARGET = bank_account_test

# Default target
all: $(TARGET)

# Build the test executable
$(TARGET): $(SOURCES) bank_account.h
	$(CXX) $(CXXFLAGS) -o $(TARGET) $(SOURCES) $(GTEST_FLAGS)

# Run the tests
test: $(TARGET)
	./$(TARGET)

# Clean build artifacts
clean:
	rm -f $(TARGET)

# Install Google Test (Ubuntu/Debian)
install-gtest:
	sudo apt-get update
	sudo apt-get install libgtest-dev cmake
	cd /usr/src/gtest && sudo cmake . && sudo make && sudo cp *.a /usr/lib

.PHONY: all test clean install-gtest