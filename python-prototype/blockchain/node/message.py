from enum import Enum
from typing import Any

import blockchain.node.socket as socket_


class MessageType(str, Enum):
    DISCOVERY = 'DISCOVERY',
    TRANSACTION = 'TRANSACTION',
    BLOCK = 'BLOCK',
    BLOCKCHAIN = 'BLOCKCHAIN',
    BLOCKCHAIN_REQUEST= 'BLOCKCHAIN_REQUEST'


class Message():

    def __init__(self, socket: socket_.Socket, message_type: MessageType, data: Any) -> None:
        self.socket = socket
        self.message_type = message_type
        self.data = data
