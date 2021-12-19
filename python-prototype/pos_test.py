import random
import string

from blockchain.proof_of_stake import ProofOfStake

if __name__ == '__main__':
    pos = ProofOfStake()
    pos.update_stake('barrie', 100)
    pos.update_stake('sjonnie', 100)
