from typing import Dict, List, Set

from .account_model import AccountModel
from .block import Block
from .transaction import Transaction, TxType
from .utils import Utils


class Blockchain():

    def __init__(self) -> None:
        self.blocks = [Block.genesis()]
        self.account_model = AccountModel()

    def to_json(self) -> Dict:
        json_dict = {}
        json_dict['blocks'] = [block.to_json() for block in self.blocks]
        return json_dict

    def latest_previous_hash(self) -> str:
        return Utils.hash(self.blocks[-1].payload()).hexdigest()

    def latest_block_height(self) -> int:
        return self.blocks[-1].height

    def is_valid_block_count(self, block: Block) -> bool:
        return self.blocks[-1].height == block.height - 1

    def is_valid_previous_block_hash(self, block: Block) -> bool:
        return self.latest_previous_hash() == block.previous_hash

    def add_block(self, block: Block) -> None:
        if self.is_valid_block_count(block) and self.is_valid_previous_block_hash(block):
            self.execute_transactions(block.transactions)
            self.blocks.append(block)

    def is_transaction_covered(self, transaction: Transaction) -> bool:
        # TODO remove echange transaction type check if type removed
        if transaction.tx_type == TxType.EXCHANGE:
            return True
        sender_balance = self.account_model.get_balance(
            transaction.sender_public_key)
        return sender_balance >= transaction.amount

    def get_covered_transactions(self, transactions: List[Transaction]) -> List[Transaction]:
        return [tx for tx in transactions if self.is_transaction_covered(tx)]

    def execute_transaction(self, transaction: Transaction) -> None:
        sender = transaction.sender_public_key
        receiver = transaction.receiver_public_key
        amount = transaction.amount
        self.account_model.update_balance(sender, -amount)
        self.account_model.update_balance(receiver, amount)

    def execute_transactions(self, transactions: List[Transaction]) -> None:
        for tx in transactions:
            self.execute_transaction(tx)
    
    def get_account_balance(self, public_key_string: str) -> int:
        return self.account_model.get_balance(public_key_string)
