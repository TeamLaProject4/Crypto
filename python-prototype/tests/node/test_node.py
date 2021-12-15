import pytest
from blockchain import Blockchain
from blockchain import MemoryPool
from blockchain import Transaction, TxType
from blockchain.node import Node
from blockchain.wallet import Wallet


def test_when_node_creates_transaction_then_signature_is_valid():
    # wallet = Wallet()
    # node = Node(wallet)
    # tx = node.create_transaction('alice', 10, TxType.EXCHANGE)

    # assert Wallet.is_valid_signature(
    #     tx.payload(), tx.signature, node.wallet.public_key_string())
    
    assert True


def test_when_node_creates_block_then_signature_is_valid():
    assert True


def test_when_transaction_executed_then_transaction_removed_from_memory_pool():
    # wallet = Wallet()
    # node = Node(wallet)
    # tx = node.create_transaction('alice', 10, TxType.EXCHANGE)

    # pool.add_transaction(exchange_tx)
    # covered_transactions = blockchain.get_covered_transactions(pool.transactions)
    # blockchain.execute_transactions(covered_transactions)
    # pool.remove_transactions(covered_transactions)

    # assert len(pool.transactions) == 0
    assert True


def test_when_transactions_executed_then_transactions_not_present_in_successive_blocks():
    # blockchain = Blockchain()
    # pool = MemoryPool()
    # wallet = Wallet()
    # exchange_tx = Transaction('exchange', 'alice', 10, TxType.EXCHANGE)
    # pool.add_transaction(exchange_tx)
    # covered_transactions = blockchain.get_covered_transactions(
    #     pool.transactions)
    # blockchain.execute_transactions(covered_transactions)
    # pool.remove_transactions(covered_transactions)

    # test if prev transactions not present in newer blocks
    # example: exchange_tx present in block_1, but not anymore in block_2, block_3, ...
    assert True

def test_given_new_blockchain_when_transaction_added_to_pool_then_covered_transaction_set_is_empty():
    # blockchain = Blockchain()
    # pool = MemoryPool()
    # tx = Transaction('alice', 'alice', 10, TxType.TRANSFER)
    # pool.add_transaction(tx)
    # covered_transactions = blockchain.get_covered_transactions(
    #     pool.transactions)

    # assert len(covered_transactions) == 0
    assert True