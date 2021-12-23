import pytest
from blockchain.wallet import WalletUtils

MNEMONIC = 'modify prepare drink head winter pizza hover trade vendor awful fruit board'

SEED_BYTES_HEX = ('f5ae3c66606928709a2f91f47bcea27870954ee0015'
                  '20667c17613f1cf1fce1ec9af963a54d4ef3997dbfe'
                  'd22459e5893d8e3d1fae949fdaf061f023634bf2d4')

PRIVATE_KEY = ('xprv9s21ZrQH143K3Q6B2KgaHL5gh4v6QkwZo3Fia8hyVKRPJC2gxNhu'
               'R5eHgHHSstL2zBAdhTiW6UoYPeqvcsw5WK7WiKQMA3Wvr5gXUDDBART')


def test_when_mnemonic_generated_then_mnemonic_is_valid():
    mnemonic = WalletUtils.generate_mnemonic()

    assert WalletUtils.is_valid_mnemonic(mnemonic)


def test_when_deterministic_mnemonic_created_then_seed_bytes_is_correct():
    assert WalletUtils.generate_seed_bytes_string(MNEMONIC) == SEED_BYTES_HEX

def test_when_deterministic_mnemonic_created_then_private_key_is_correct():
    seed_bytes = WalletUtils.generate_seed_bytes(MNEMONIC)

    assert WalletUtils.construct_private_key(seed_bytes) == PRIVATE_KEY
