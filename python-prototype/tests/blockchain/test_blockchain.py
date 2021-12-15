import pytest
from blockchain import Block
from blockchain import Blockchain
from blockchain import Transaction, TxType


def test_when_new_blockchain_created_then_first_block_is_genesis():
    blockchain = Blockchain()
    genesis = Block.genesis()

    assert blockchain.blocks[0] == genesis


def test_when_new_block_has_correct_previous_hash_then_previous_hash_valid():
    blockchain = Blockchain()
    prev_hash = blockchain.latest_previous_hash()
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    block = Block([tx], prev_hash, 'forger', 1)

    assert blockchain.is_valid_previous_block_hash(block)


def test_when_new_block_added_then_latest_block_height_is_correct():
    blockchain = Blockchain()
    prev_hash = blockchain.latest_previous_hash()
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    block = Block([tx], prev_hash, 'forger', 1)
    blockchain.add_block(block)

    assert blockchain.latest_block_height() == 1


def test_given_new_blockchain_when_block_added_then_second_block_is_block():
    blockchain = Blockchain()
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    prev_hash = blockchain.latest_previous_hash()
    block = Block([tx], prev_hash, 'forger', 1)
    blockchain.add_block(block)

    assert blockchain.blocks[1] == block


def test_when_third_block_has_lower_block_count_then_block_count_invalid():
    blockchain = Blockchain()
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    prev_hash_1 = blockchain.latest_previous_hash()
    block_1 = Block([tx], prev_hash_1, 'forger', 1)
    blockchain.add_block(block_1)
    prev_hash_2 = blockchain.latest_previous_hash()
    block_2 = Block([tx], prev_hash_2, 'forger', 1)

    assert not blockchain.is_valid_block_count(block_2)


def test_when_third_block_has_higher_block_count_then_block_count_invalid():
    blockchain = Blockchain()
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    prev_hash_1 = blockchain.latest_previous_hash()
    block_1 = Block([tx], prev_hash_1, 'forger', 1)
    blockchain.add_block(block_1)
    prev_hash_2 = blockchain.latest_previous_hash()
    block_2 = Block([tx], prev_hash_2, 'forger', 3)

    assert not blockchain.is_valid_block_count(block_2)


def test_when_third_block_has_correct_block_count_then_block_count_valid():
    blockchain = Blockchain()
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    prev_hash_1 = blockchain.latest_previous_hash()
    block_1 = Block([tx], prev_hash_1, 'forger', 1)
    blockchain.add_block(block_1)
    prev_hash_2 = blockchain.latest_previous_hash()
    block_2 = Block([tx], prev_hash_2, 'forger', 2)

    assert blockchain.is_valid_block_count(block_2)


def test_given_new_blockchain__when_transaction_created_then_transaction_not_covered():
    blockchain = Blockchain()
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)

    assert not blockchain.is_transaction_covered(tx)


# TODO remove test when exchange transaction type removed
def test_when_exchange_transaction_then_transaction_is_covered():
    blockchain = Blockchain()
    tx = Transaction('alice', 'bob', 10, TxType.EXCHANGE)

    assert blockchain.is_transaction_covered(tx)


# TODO remove test when exchange transaction type removed
def test_when_exchange_transaction_executed_then_receiver_has_correct_balance():
    blockchain = Blockchain()
    exchange_tx = Transaction('exchange', 'alice', 10, TxType.EXCHANGE)
    covered_transactions = blockchain.get_covered_transactions([exchange_tx])
    blockchain.execute_transactions(covered_transactions)

    assert blockchain.get_account_balance('alice') == 10


def test_when_transfer_transaction_executed_then_receiver_has_correct_balance():
    blockchain = Blockchain()
    exchange_tx = Transaction('exchange', 'alice', 10, TxType.EXCHANGE)
    covered_transactions = blockchain.get_covered_transactions([exchange_tx])
    blockchain.execute_transactions(covered_transactions)
    transfer_tx = Transaction('alice', 'bob', 5, TxType.TRANSFER)
    covered_transactions = blockchain.get_covered_transactions([transfer_tx])
    blockchain.execute_transactions(covered_transactions)

    assert blockchain.get_account_balance('bob') == 5
