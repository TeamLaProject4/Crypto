from typing import Dict, List

from blockchain.proof_of_stake import ProofOfStake

from .account_model import AccountModel
from .block import Block
from .transaction import Transaction, TxType
from .utils import BlockchainUtils


class Blockchain():

    def __init__(self) -> None:
        self.blocks = [Block.genesis()]
        self.account_model = AccountModel()
        self.pos = ProofOfStake()

    def to_json(self) -> Dict:
        json_dict = {}
        json_dict['blocks'] = [block.to_json() for block in self.blocks]
        return json_dict

    def latest_block_height(self) -> int:
        return self.blocks[-1].height

    def add_block(self, block: Block) -> None:
        if self.is_valid_block_count(block) and self.is_valid_previous_block_hash(block):
            self.execute_transactions(block.transactions)
            self.blocks.append(block)

    def execute_transactions(self, transactions: List[Transaction]) -> None:
        for tx in transactions:
            self.execute_transaction(tx)

    # TODO: unit test for stake transactions
    # TODO: function too long: seperate responsibilities
    def execute_transaction(self, transaction: Transaction) -> None:
        if transaction.tx_type == TxType.STAKE:
            sender = transaction.sender_public_key
            receiver = transaction.receiver_public_key
            if sender == receiver:
                amount = transaction.amount
                self.pos.update_stake(sender, amount)
                self.account_model.update_balance(sender, -amount)
        else:
            sender = transaction.sender_public_key
            receiver = transaction.receiver_public_key
            amount = transaction.amount
            self.account_model.update_balance(sender, -amount)
            self.account_model.update_balance(receiver, amount)

    def is_valid_block_count(self, block: Block) -> bool:
        return self.blocks[-1].height == block.height - 1

    def is_valid_previous_block_hash(self, block: Block) -> bool:
        return self.latest_previous_hash() == block.previous_hash
    
    # TODO: unit test
    def is_valid_forger(self, block: Block) -> bool:
        return block.forger == self.next_forger()

    # TODO: unit test
    def is_block_transactions_valid(self, block: Block) -> bool:
        transactions = block.transactions
        covered_transactions = self.get_covered_transactions(transactions)
        return len(transactions) == len(covered_transactions)

    def latest_previous_hash(self) -> str:
        return BlockchainUtils.hash(self.blocks[-1].payload()).hexdigest()

    def get_covered_transactions(self, transactions: List[Transaction]) -> List[Transaction]:
        return [tx for tx in transactions if self.is_transaction_covered(tx)]

    def is_transaction_covered(self, transaction: Transaction) -> bool:
        # TODO: remove echange transaction type check if type removed
        # TODO: check if double spending is covered (unit test: amount == 20, 3 transactions all send 10)
        if transaction.tx_type == TxType.EXCHANGE:
            return True
        sender_balance = self.account_model.get_balance(
            transaction.sender_public_key)
        return sender_balance >= transaction.amount

    def get_account_balance(self, public_key_string: str) -> int:
        return self.account_model.get_balance(public_key_string)

    # TODO: unit test
    def next_forger(self) -> str:
        prev_block_hash = self.latest_previous_hash()
        return self.pos.pick_forger(prev_block_hash)

    # TODO: unit test
    # TODO: cleanup function + add type hinting
    def create_block(self, transactions: List[Transaction], forger_wallet) -> Block:
        covered_transactions = self.get_covered_transactions(
            transactions)
        self.execute_transactions(covered_transactions)
        block = forger_wallet.create_block(
            covered_transactions, self.latest_previous_hash(), self.latest_block_height() + 1)
        self.add_block(block)
        return block
    
    # TODO: unit test
    # TODO: optimize search
    def is_transaction_in_blockchain(self, transaction: Transaction) -> bool:
        for block in self.blocks:
            for block_tx in block.transactions:
                if block_tx == transaction:
                    return True
        return False
