import blockchain.blockchain.utils as utils


class Lot():

    def __init__(self, public_key_string: str, iteration: int, previous_block_hash: str) -> None:
        self.public_key_string = public_key_string
        self.iteration = iteration
        self.previous_block_hash = previous_block_hash

    def hash(self) -> str:
        hash = self.public_key_string + self.previous_block_hash
        for _ in range(self.iteration):
            hash = utils.BlockchainUtils.hash(hash).hexdigest()
        return hash
