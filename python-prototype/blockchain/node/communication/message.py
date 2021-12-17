from enum import Enum

import blockchain.node.communication.message as message
import blockchain.node.communication.socket as socket


class Message():

    def __init__(self, socket: socket.Socket, msg_type: message.MessageType, data) -> None:
        self.sender_connector = socket
        self.msg_type = msg_type
        self.data = data


class MessageType(str, Enum):
    DISCOVERY = 'DISCOVERY'
