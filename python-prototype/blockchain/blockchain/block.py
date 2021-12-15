from __future__ import annotations

import time
from copy import deepcopy
from typing import Dict, List

from .transaction import Transaction


class Block():

    def __init__(self,
                 transactions: List[Transaction],
                 previous_hash: str,
                 forger: str,
                 height: int,
                 timestamp: int = None,
                 signature: str = None
                 ) -> None:
        self.transactions = transactions
        self.previous_hash = previous_hash
        self.forger = forger
        self.height = height
        self.timestamp = timestamp if timestamp != None else int(time.time())
        self.signature = signature if signature != None else ''

    # TODO rewrite method as: (self, __o: object)
    def __eq__(self, block: Block) -> bool:
        return (
            self.signature == block.signature
            and self.previous_hash == block.previous_hash
            and self.forger == block.forger
            and self.height == block.height
            and self.timestamp == block.timestamp
        )

    def to_json(self) -> Dict:
        json_dict = {}
        json_dict['previous_hash'] = self.previous_hash
        json_dict['forger'] = self.forger
        json_dict['height'] = self.height
        json_dict['timestamp'] = self.timestamp
        json_dict['signature'] = self.signature
        json_dict['transactions'] = [tx.to_json() for tx in self.transactions]
        return json_dict

    # TODO: Find out how solution below is possible
    # def to_json(self) -> Dict:
    #     json_dict = self.__dict__
    #     json_dict['transactions'] = [tx.to_json() for tx in self.transactions]
    #     return json_dict

    def payload(self) -> Dict:
        json_repr = deepcopy(self.to_json())
        json_repr['signature'] = ''
        return json_repr

    def sign(self, signature: str) -> None:
        self.signature = signature

    @staticmethod
    def genesis() -> Block:
        genesis = Block([], 'genesis_hash',
                        'genesis_forger', 0, timestamp=0)
        return genesis
