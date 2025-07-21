class BankAccount:
    """Simple checking account with deposit and withdraw."""
    def __init__(self, name: str, balance: float = 0.0):
        self._name = name
        self._balance = balance

    def deposit(self, amount: float) -> None:
        if amount <= 0:
            raise ValueError("Deposit must be positive")
        self._balance += amount

    def withdraw(self, amount: float) -> float:
        if amount <= 0:
            raise ValueError("Withdraw must be positive")
        if amount > self._balance:
            raise ValueError("Insufficient funds")
        self._balance -= amount
        return amount

    def __repr__(self) -> str:
        return f"{self._name}: ${self._balance:.2f}"