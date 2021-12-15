import time
import uuid

import pytest
from blockchain import Transaction, TxType

SENDER_PUBLIC_KEY = 'alice'
RECEIVER_PUBLIC_KEY = 'bob'
AMOUNT = 10
TX_TYPE = TxType.TRANSFER
SIGNATURE = 'VeRyN1CeSiGn4TuRe'


def create_transaction() -> Transaction:
    return Transaction(SENDER_PUBLIC_KEY, RECEIVER_PUBLIC_KEY, AMOUNT, TX_TYPE)


def test_when_new_transaction_constructed_then_signature_is_empty():
    # TODO: Find out how to prevent duplicate code here (aka, how to run setup code with pytest)
    tx = create_transaction()

    assert tx.signature == ''


def test_when_duplicate_transactions_constructed_then_transactions_are_equal():
    id = uuid.uuid4().hex
    timestamp = int(time.time())
    tx1 = Transaction(SENDER_PUBLIC_KEY, RECEIVER_PUBLIC_KEY,
                      AMOUNT, TX_TYPE, id=id, timestamp=timestamp)
    tx2 = Transaction(SENDER_PUBLIC_KEY, RECEIVER_PUBLIC_KEY,
                      AMOUNT, TX_TYPE, id=id, timestamp=timestamp)

    assert tx1 == tx2


def test_when_transaction_signed_then_signature_is_set():
    tx = create_transaction()
    tx.sign(SIGNATURE)

    assert tx.signature == SIGNATURE


def test_when_transaction_signed_then_payload_signature_stays_empty():
    tx = create_transaction()
    tx.sign(SIGNATURE)
    payload = tx.payload()

    assert payload['signature'] == ''
