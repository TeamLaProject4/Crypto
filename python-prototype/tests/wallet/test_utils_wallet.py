import pytest
from blockchain.wallet import WalletUtils

WORDS = ['modify', 'prepare', 'drink', 'head', 'winter', 'pizza',
         'hover', 'trade', 'vendor', 'awful', 'fruit', 'board']

SEED_BYTES_HEX = ('f5ae3c66606928709a2f91f47bcea27870954ee0015'
                  '20667c17613f1cf1fce1ec9af963a54d4ef3997dbfe'
                  'd22459e5893d8e3d1fae949fdaf061f023634bf2d4')


def test_when_deterministic_mnemonic_created_then_words_list_is_correct():
    mnemonic = WalletUtils.generate_mnemonic_from_words(WORDS)

    assert WalletUtils.get_mnemonic_words(mnemonic) == WORDS


def test_when_deterministic_mnemonic_created_then_seed_bytes_is_correct():
    mnemonic = WalletUtils.generate_mnemonic_from_words(WORDS)

    assert WalletUtils.generate_seed_bytes_string(mnemonic) == SEED_BYTES_HEX
