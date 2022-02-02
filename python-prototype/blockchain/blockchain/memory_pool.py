from typing import List

from .transaction import Transaction

TRANSACTION_THRESHOLD = 100
TRANSACTION_THRESHOLD = 100


class MemoryPool():

    def __init__(self, transaction_threshold = TRANSACTION_THRESHOLD) -> None:
        self.transactions = []
        self.transaction_threshold = transaction_threshold

    def is_transaction_in_pool(self, transaction: Transaction) -> bool:
        return transaction in self.transactions

    def add_transaction(self, transaction: Transaction) -> None:
        if transaction not in self.transactions:
            self.transactions.append(transaction)

    def add_transactions(self, transactions: List[Transaction]) -> None:
        for tx in transactions:
            self.add_transaction(tx)

    def remove_transaction(self, transaction: Transaction) -> None:
        self.transactions.remove(transaction)

    def remove_transactions(self, transactions: List[Transaction]) -> None:
        for tx in transactions:
            self.remove_transaction(tx)

    # TODO: unit test
    def is_transaction_threshold_reached(self) -> bool:
        return len(self.transactions) >= self.transaction_threshold
