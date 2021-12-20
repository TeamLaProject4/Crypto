from typing import Any

import blockchain.node.communication as communication
import blockchain.node.message as message_
import blockchain.node.rest_api as rest_api
from blockchain import Block, Blockchain, MemoryPool, Transaction
from blockchain.node.encoding import Encoding
from blockchain.wallet import Wallet


class Node():

    def __init__(self,
                 host: str,
                 port: int,
                 wallet: Wallet,
                 blockchain: Blockchain = None,
                 memory_pool: MemoryPool = None,
                 ) -> None:
        self.host = host
        self.port = port
        self.wallet = wallet
        self.blockchain = blockchain if blockchain != None else Blockchain()
        self.memory_pool = memory_pool if memory_pool != None else MemoryPool()

    def start_p2p(self) -> None:
        self.p2p = communication.Communication(
            self.host, self.port)
        self.p2p.start(self)

    def start_rest_api(self, port: int) -> None:
        self.api = rest_api.RestAPI()
        self.api.inject_node(self)
        self.api.start(port)

    def is_valid_signature(self, transaction_payload: Any, signature: str, sender_public_key: str) -> bool:
        return Wallet.is_valid_signature(transaction_payload, signature, sender_public_key)

    def is_transaction_in_memory_pool(self, transaction: Transaction) -> bool:
        return self.memory_pool.is_transaction_in_pool(transaction)

    # TODO: function too long: seperate responsibilities
    def handle_transaction(self, transaction: Transaction) -> None:
        payload = transaction.payload()
        signature = transaction.signature
        sender_public_key = transaction.sender_public_key
        transaction_in_memory_pool = self.is_transaction_in_memory_pool(
            transaction)
        signature_valid = self.is_valid_signature(
            payload, signature, sender_public_key)
        transaction_in_blockchain = self.blockchain.is_transaction_in_blockchain(
            transaction)

        if not transaction_in_memory_pool and signature_valid and not transaction_in_blockchain:
            self.memory_pool.add_transaction(transaction)
            message = message_.Message(
                self.p2p.socket, message_.MessageType.TRANSACTION, transaction)
            encoded_message = Encoding.encode(message)
            self.p2p.broadcast_message(encoded_message)
            if self.memory_pool.is_transaction_threshold_reached():
                self.forge()

    def handle_block(self, block: Block) -> None:
        payload = block.payload()
        signature = block.signature
        forger_public_key = block.forger

        block_count_valid = self.blockchain.is_valid_block_count(block)
        previous_block_hash_valid = self.blockchain.is_valid_previous_block_hash(block)
        signature_valid = self.is_valid_signature(payload, signature, forger_public_key)
        forger_valid = self.blockchain.is_valid_forger(block)
        block_transactions_valid = self.blockchain.is_block_transactions_valid(block)

        if all([block_count_valid, previous_block_hash_valid, signature_valid, forger_valid, block_transactions_valid]):
            self.blockchain.add_block(block)
            self.memory_pool.remove_transactions(block.transactions)
            message = message_.Message(self.p2p.socket, message_.MessageType.BLOCK, block)
            encoded_message = Encoding.encode(message)
            self.p2p.broadcast_message(encoded_message)

    # TODO: unit test
    # TODO: function too long: seperate responsibilities
    def forge(self) -> None:
        forger = self.blockchain.next_forger()
        if forger == self.wallet.public_key_string():
            print('Is next forger')
            transactions = self.memory_pool.transactions
            block = self.blockchain.create_block(transactions, self.wallet)
            self.memory_pool.remove_transactions(block.transactions)
            message = message_.Message(
                self.p2p.socket, message_.MessageType.BLOCK, block)
            encoded_message = Encoding.encode(message)
            self.p2p.broadcast_message(encoded_message)
        else:
            print('Not next forger')
