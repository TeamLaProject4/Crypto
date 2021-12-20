from typing import List

import blockchain.blockchain.utils as utils
from blockchain.proof_of_stake import Lot


class ProofOfStake():

    def __init__(self) -> None:
        self.stakers = {}
        self.set_genesis_node_stake()
    
    # TODO: unit test
    # TODO: remove hardcoded path and add more generic solution
    # TODO: after others stake, then genesis can no longer be forger, if everyone unstaked, then genesis can be forger
    def set_genesis_node_stake(self) -> None:
        with open('keys/genesisPublicKey.pem', 'r') as f:
            genesis_public_key = f.read()
        self.stakers[genesis_public_key] = 1

    def is_account_in_stakers(self, public_key_string: str) -> bool:
        return public_key_string in self.stakers.keys()

    def add_account_to_stakers(self, public_key_string: str) -> None:
        if not self.is_account_in_stakers(public_key_string):
            self.stakers[public_key_string] = 0

    def update_stake(self, public_key_string: str, stake: int):
        # add account if not present in stakers
        self.add_account_to_stakers(public_key_string)
        self.stakers[public_key_string] += stake

    def get_stake(self, public_key_string: str) -> int:
        # add account if not present in stakers
        self.add_account_to_stakers(public_key_string)
        return self.stakers[public_key_string]

    def __generate_lots(self, seed: str) -> List[Lot]:

        def generate_staker_lots(public_key_string: str) -> List[Lot]:
            stake_amount = self.get_stake(public_key_string)
            return [Lot(public_key_string, stake + 1, seed) for stake in range(stake_amount)]

        lots = []
        for staker in self.stakers.keys():
            lots += generate_staker_lots(staker)
        return lots

    def __pick_winner(self, lots: List[Lot], seed: str) -> Lot:

        def is_offset_smaller(lot: Lot) -> bool:
            offset = get_offset(lot)
            return offset < least_offset

        def get_offset(lot: Lot) -> int:
            lot_int = int(lot.hash(), 16)
            return abs(lot_int - seed_int)

        winner = lots[0]
        least_offset = float('INF')
        seed_int = int(
            utils.BlockchainUtils.hash(seed).hexdigest(), 16)

        for lot in lots:
            if is_offset_smaller(lot):
                least_offset = get_offset(lot)
                winner = lot

        return winner

    def pick_forger(self, previous_block_hash: str) -> str:
        lots = self.__generate_lots(previous_block_hash)
        winner = self.__pick_winner(lots, previous_block_hash)
        return winner.public_key_string
