from blockchain import Blockchain, MemoryPool, Transaction, TxType
from blockchain.wallet import Wallet

from .communication import SocketCommunication


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
    
    def start(self) -> None:
        self.p2p = SocketCommunication(self.host, self.port)
        self.p2p.start()

    # def create_transaction(self, receiver_public_key: str, amount: int, tx_type: TxType) -> Transaction:
    #     tx = self.wallet.create_transaction(receiver_public_key, amount, tx_type)
    #     self.memory_pool.add_transaction(tx)
    #     return tx

    # def execute_transaction(self, transaction: Transaction) -> None:
    #     self.blockchain.execute_transaction(transaction)
    #     self.memory_pool.remove_transaction(transaction)
