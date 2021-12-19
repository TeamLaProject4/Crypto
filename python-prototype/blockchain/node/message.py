from enum import Enum

import blockchain.node.socket as socket


class MessageType(str, Enum):
    DISCOVERY = 'DISCOVERY'


class Message():

    def __init__(self, socket: socket.Socket, message_type: MessageType, data) -> None:
        self.socket = socket
        self.message_type = message_type
        self.data = data
