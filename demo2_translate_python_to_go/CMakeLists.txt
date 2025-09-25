# [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-09-17
cmake_minimum_required(VERSION 3.14)
project(BankAccountTest)

# Set C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Find required packages
find_package(GTest REQUIRED)

# Include directories
include_directories(${GTEST_INCLUDE_DIRS})

# Create executable for tests
add_executable(bank_account_test
    bank_account.cc
    bank_account_test.cc
)

# Link libraries
target_link_libraries(bank_account_test
    ${GTEST_LIBRARIES}
    ${GTEST_MAIN_LIBRARIES}
    pthread
)

# Enable testing
enable_testing()

# Add test
add_test(NAME BankAccountUnitTests COMMAND bank_account_test)