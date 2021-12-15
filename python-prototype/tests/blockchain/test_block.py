import time

import pytest
from blockchain import Block
from blockchain import Transaction, TxType


def create_block() -> Block:
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    return Block([tx], 'prev_hash', 'forger', 1)


def test_when_new_block_constructed_then_signature_is_empty():
    # TODO: Find out how to prevent duplicate code here (aka, how to run setup code with pytest)
    block = create_block()

    assert block.signature == ''


def test_when_duplicate_blocks_constructed_then_blocks_are_equal():
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    timestamp = int(time.time())
    block1 = Block([tx], 'prev_hash', 'forger', 1, timestamp=timestamp)
    block2 = Block([tx], 'prev_hash', 'forger', 1, timestamp=timestamp)

    assert block1 == block2


def test_when_block_signed_then_signature_is_set():
    block = create_block()
    signature = 'signature'
    block.sign(signature)

    assert block.signature == signature


def test_when_block_signed_then_payload_signature_stays_empty():
    block = create_block()
    signature = 'signature'
    payload = block.payload()

    assert payload['signature'] == ''
