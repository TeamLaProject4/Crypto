from typing import Any

import blockchain.node.communication as communication
import blockchain.node.message as message_
import blockchain.node.rest_api as rest_api
from blockchain import Blockchain, MemoryPool, Transaction
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

    def handle_transaction(self, transaction: Transaction) -> None:
        payload = transaction.payload()
        signature = transaction.signature
        sender_public_key = transaction.sender_public_key

        if not self.is_transaction_in_memory_pool(transaction) and self.is_valid_signature(payload, signature, sender_public_key):
            self.memory_pool.add_transaction(transaction)
            message = message_.Message(
                self.p2p.socket, message_.MessageType.TRANSACTION, transaction)
            encoded_message = Encoding.encode(message)
            self.p2p.broadcast_message(encoded_message)
