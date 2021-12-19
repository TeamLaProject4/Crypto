import random
import string

import pytest
from blockchain.proof_of_stake import ProofOfStake


def generate_random_string(length: int) -> str:
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters) for _ in range(length))


def test_when_new_pos_then_account_not_in_stakers():
    pos = ProofOfStake()

    assert pos.is_account_in_stakers('satoshi') == False


def test_when_account_added_then_account_in_stakers():
    pos = ProofOfStake()
    pos.add_account_to_stakers('barrie')

    assert pos.is_account_in_stakers('barrie') == True


def test_when_account_added_then_stake_of_account_is_zero():
    pos = ProofOfStake()
    pos.add_account_to_stakers('barrie')

    assert pos.get_stake('barrie') == 0


def test_when_stake_updated_then_stake_of_account_updated():
    pos = ProofOfStake()
    pos.add_account_to_stakers('barrie')
    pos.update_stake('barrie', 50)

    assert pos.get_stake('barrie') == 50


def test_when_one_staker_added_then_staker_is_forger():
    pos = ProofOfStake()
    pos.update_stake('barrie', 100)

    assert pos.pick_forger('prev_hash') == 'barrie'


def test_when_forger_picked_then_result_is_deterministic():
    pos = ProofOfStake()
    pos.update_stake('barrie', 100)
    pos.update_stake('sjonnie', 100)
    seed = 'zaadje'
    forger = pos.pick_forger(seed)

    for _ in range(20):
        assert pos.pick_forger(seed) == forger


def test_given_equal_stake_when_forger_picked_then_it_represents_stake():
    pos = ProofOfStake()
    pos.update_stake('barrie', 100)
    pos.update_stake('sjonnie', 100)
    wins = {'barrie': 0, 'sjonnie': 0}

    for _ in range(100):
        seed = generate_random_string(16)
        forger = pos.pick_forger(seed)
        wins[forger] += 1
    
    b_wins = wins['barrie']
    s_wins = wins['sjonnie']
    assert 40 < b_wins < 60 and 40 < s_wins < 60

def test_given_unequal_stake_when_forger_picked_then_it_represents_stake():
    pos = ProofOfStake()
    pos.update_stake('barrie', 10)
    pos.update_stake('sjonnie', 90)
    wins = {'barrie': 0, 'sjonnie': 0}

    for _ in range(100):
        seed = generate_random_string(16)
        forger = pos.pick_forger(seed)
        wins[forger] += 1
    
    b_wins = wins['barrie']
    s_wins = wins['sjonnie']
    assert 0 < b_wins < 20 and 80 < s_wins < 100
