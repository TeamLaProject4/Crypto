from __future__ import annotations

import time
import uuid
from copy import deepcopy
from enum import Enum
from typing import Dict


class Transaction():

    def __init__(self,
                 sender_public_key: str,
                 receiver_public_key: str,
                 amount: int,
                 tx_type: TxType,
                 id: str = None,
                 timestamp: int = None,
                 signature: str = None
                 ) -> None:
        self.sender_public_key = sender_public_key
        self.receiver_public_key = receiver_public_key
        self.amount = amount
        self.tx_type = tx_type
        self.id = id if id != None else uuid.uuid4().hex
        self.timestamp = timestamp if timestamp != None else int(time.time())
        self.signature = signature if signature != None else ''

    # TODO rewrite method as: (self, __o: object)
    def __eq__(self, transaction: Transaction) -> bool:
        return self.id == transaction.id
    
    def __hash__(self):
        return hash(repr(self))

    def to_json(self) -> Dict:
        return self.__dict__

    def payload(self) -> Dict:
        json_repr = deepcopy(self.to_json())
        json_repr['signature'] = ''
        return json_repr

    def sign(self, signature: str) -> None:
        self.signature = signature


# TODO remove exchange transaction type when initial coins has been implemented
class TxType(str, Enum):
    TRANSFER = 'TRANSFER',
    EXCHANGE = 'EXCHANGE',
    STAKE = 'STAKE'
