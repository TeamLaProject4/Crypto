import pytest

from blockchain import Transaction, TxType
from blockchain import MemoryPool


def test_when_transaction_not_added_then_transaction_not_in_pool():
    pool = MemoryPool()
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)

    assert not pool.is_transaction_in_pool(tx)


def test_when_transaction_added_then_transaction_in_pool():
    pool = MemoryPool()
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    pool.add_transaction(tx)

    assert pool.is_transaction_in_pool(tx)


def test_when_multiple_transactions_added_then_transactions_in_pool():
    pool = MemoryPool()
    tx_1 = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    tx_2 = Transaction('bob', 'alice', 5, TxType.TRANSFER)
    pool.add_transactions([tx_1, tx_2])

    assert all([pool.is_transaction_in_pool(tx_1),
               pool.is_transaction_in_pool(tx_2)])


def test_when_transaction_added_twice_then_one_transaction_in_pool():
    pool = MemoryPool()
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    pool.add_transaction(tx)
    pool.add_transaction(tx)

    assert len(pool.transactions) == 1


def test_given_pool_with_transaction_when_transaction_removed_transaction_then_transaction_not_in_pool():
    pool = MemoryPool()
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    pool.add_transaction(tx)
    pool.remove_transaction(tx)

    assert not pool.is_transaction_in_pool(tx)


def test_given_pool_with_transactions_when_transactions_removed_from_pool_then_transactions_not_in_pool():
    pool = MemoryPool()
    tx_1 = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    tx_2 = Transaction('bob', 'alice', 5, TxType.TRANSFER)
    tx_3 = Transaction('alice', 'bob', 1, TxType.TRANSFER)
    pool.add_transactions([tx_1, tx_2, tx_3])
    pool.remove_transactions([tx_1, tx_3])

    assert not all([pool.is_transaction_in_pool(tx_1),
                   pool.is_transaction_in_pool(tx_3)])


def test_given_pool_with_three_transactions_when_last_two_transactions_removed_then_first_transaction_in_pool():
    pool = MemoryPool()
    tx_1 = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    tx_2 = Transaction('bob', 'alice', 5, TxType.TRANSFER)
    tx_3 = Transaction('alice', 'bob', 1, TxType.TRANSFER)
    pool.add_transactions([tx_1, tx_2, tx_3])
    pool.remove_transactions([tx_2, tx_3])

    assert pool.is_transaction_in_pool(tx_1)
