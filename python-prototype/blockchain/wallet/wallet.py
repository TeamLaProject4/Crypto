from typing import Any, List

from Crypto.PublicKey import RSA
from Crypto.Signature import PKCS1_v1_5

from blockchain import Block
from blockchain import Transaction, TxType
from blockchain import BlockchainUtils

KEY_LENGTH_BITS = 2048


class Wallet():

    def __init__(
            self,
            private_key: str = None
    ) -> None:
        self.key_pair = RSA.generate(
            KEY_LENGTH_BITS) if private_key == None else RSA.import_key(private_key)

    def sign(self, data: Any) -> Any:
        data_hash = BlockchainUtils.hash(data)
        signature_scheme_object = PKCS1_v1_5.new(self.key_pair)
        signature = signature_scheme_object.sign(data_hash)
        return signature.hex()

    def public_key_string(self) -> str:
        return self.key_pair.public_key().export_key('PEM').decode('utf-8')

    def create_transaction(self, receiver_public_key: str, amount: int, tx_type: TxType) -> Transaction:
        transaction = Transaction(
            self.public_key_string(), receiver_public_key, amount, tx_type)
        signature = self.sign(transaction.to_json())
        transaction.sign(signature)
        return transaction

    def create_block(self, transactions: List[Transaction], previous_hash: str, block_count: int) -> Block:
        block = Block(transactions, previous_hash,
                      self.public_key_string(), block_count)
        signature = self.sign(block.payload())
        block.sign(signature)
        return block

    @staticmethod
    def is_valid_signature(data: Any, signature: str, public_key_string: str) -> bool:
        signature = bytes.fromhex(signature)
        data_hash = BlockchainUtils.hash(data)
        public_key = RSA.importKey(public_key_string)
        signature_scheme_object = PKCS1_v1_5.new(public_key)
        return signature_scheme_object.verify(data_hash, signature)
