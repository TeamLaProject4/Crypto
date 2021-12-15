from typing import List

from .transaction import Transaction


class MemoryPool():

    def __init__(self) -> None:
        self.transactions = []

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
