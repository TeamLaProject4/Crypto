import pytest
from blockchain.proof_of_stake import Lot


def test_when_lot_hashed_then_hash_is_correct():
    lot = Lot('alice', 1, 'prev_hash')

    assert lot.hash() == 'e61505c60c79e5ae93f3b7d943f8b76e58e3e87524b4fa06ee279bc9e9142049'
