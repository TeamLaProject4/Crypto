from enum import Enum

import blockchain.node.socket as socket


class MessageType(str, Enum):
    DISCOVERY = 'DISCOVERY'


class Message():

    def __init__(self, socket: socket.Socket, msg_type: MessageType, data) -> None:
        self.sender_connector = socket
        self.msg_type = msg_type
        self.data = data
