from threading import Thread
import time

from p2pnetwork.node import Node

from blockchain.node.communication.message import Message, MessageType
from blockchain.node.communication.encoding import Encoding


class PeerDiscoveryHandler():

    def __init__(self, node) -> None:
        self.socket_communication = node

    def start(self) -> None:
        thread_status = Thread(target=self.status, args=())
        thread_discovery = Thread(target=self.discovery, args=())
        thread_status.start()
        thread_discovery.start()

    def status(self) -> None:
        while True:
            print('status')
            time.sleep(20)

    def discovery(self) -> None:
        while True:
            print('discovery')
            time.sleep(20)

    def handshake_message(self) -> str:
        socket = self.socket_communication.socket
        peers = self.socket_communication.peers
        message_type = MessageType.DISCOVERY
        message = Message(socket, message_type, peers)
        encoded_message = Encoding.encode(message)
        return encoded_message

    def handshake(self, node: Node) -> None:
        handshake_message = self.handshake_message()
        self.socket_communication.send_message(node, handshake_message)
