import pytest
from blockchain import Transaction, TxType
from blockchain.wallet import Wallet
from Crypto.PublicKey import RSA


def test_when_transaction_signed_using_wallet_then_signature_is_set():
    wallet = Wallet()
    tx = Transaction('alice', 'bob', 10, TxType.TRANSFER)
    signature = wallet.sign(tx.payload())
    tx.sign(signature)

    assert tx.signature == signature


def test_when_wallet_set_with_key_pair_then_public_key_string_equals_public_key_from_pair():
    key_pair = RSA.generate(2048)
    private_key = key_pair.export_key('PEM').decode('utf-8')
    public_key = key_pair.public_key().export_key('PEM').decode('utf-8')
    wallet = Wallet(private_key)

    assert wallet.public_key_string() == public_key


def test_given_transaction_signed_by_wallet_when_signature_validated_using_wallet_public_key_then_its_valid():
    wallet = Wallet()
    tx = wallet.create_transaction('bob', 10, TxType.TRANSFER)

    assert Wallet.is_valid_signature(
        tx.payload(), tx.signature, wallet.public_key_string())


def test_given_transaction_signed_by_wallet_when_signature_validated_using_other_wallet_public_key_then_its_invalid():
    wallet_alice = Wallet()
    wallet_bob = Wallet()
    tx = wallet_alice.create_transaction('bob', 10, TxType.TRANSFER)

    assert not Wallet.is_valid_signature(
        tx.payload(), tx.signature, wallet_bob.public_key_string())


def test_given_block_signed_by_wallet_when_signature_validated_using_wallet_public_key_then_its_valid():
    wallet = Wallet()
    tx = wallet.create_transaction('bob', 10, TxType.TRANSFER)
    block = wallet.create_block([tx], 'prev_hash', 5)

    assert Wallet.is_valid_signature(
        block.payload(), block.signature, wallet.public_key_string())


def test_given_block_signed_by_wallet_when_signature_validated_using_other_wallet_public_key_then_its_invalid():
    wallet_alice = Wallet()
    wallet_bob = Wallet()
    tx = wallet_alice.create_transaction('bob', 10, TxType.TRANSFER)
    block = wallet_alice.create_block([tx], 'prev_hash', 5)

    assert not Wallet.is_valid_signature(
        block.payload(), block.signature, wallet_bob.public_key_string())
