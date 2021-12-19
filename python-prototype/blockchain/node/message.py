from enum import Enum
from typing import Any

import blockchain.node.socket as socket


class MessageType(str, Enum):
    DISCOVERY = 'DISCOVERY'


class Message():

    def __init__(self, socket: socket.Socket, message_type: MessageType, data: Any) -> None:
        self.socket = socket
        self.message_type = message_type
        self.data = data
