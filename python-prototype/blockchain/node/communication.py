from __future__ import annotations

import json
from typing import List

import blockchain.node.message as message_
import blockchain.node.node as node_
import blockchain.node.peer_discovery_handler as peer_discovery_handler
import blockchain.node.socket as socket_
from blockchain.node.encoding import Encoding
from p2pnetwork.node import Node

# TODO: refactor bootnode. Bootnodes should be in a config file or so and must be derived to a list
BOOTNODE = socket_.Socket('localhost', 10000)


class Communication(Node):

    # TODO: remove hardcoded BOOTNODE
    def __init__(self, host: str, port: int, bootnodes: List[socket_.Socket] = [BOOTNODE]) -> None:
        super().__init__(host, port, callback=None)
        self.peers = []
        self.peer_discovery_handler = peer_discovery_handler.PeerDiscoveryHandler(
            self)
        self.socket = socket_.Socket(host, port)
        self.bootnodes = bootnodes

    def start(self, node: node_.Node) -> None:
        super().start()
        self.node = node
        self.peer_discovery_handler.start()
        self.connect_with_bootnode()

    def connect_with_bootnode(self) -> None:
        # TODO: remove hardcoded indice 0, instead try all bootnodes, one by one till connected
        bootnode = self.bootnodes[0]
        if self.socket != bootnode:
            self.connect_with_node(bootnode.host, bootnode.port)

    def inbound_node_connected(self, node: Communication) -> None:
        self.peer_discovery_handler.handshake(node)

    def outbound_node_connected(self, node: Communication) -> None:
        self.peer_discovery_handler.handshake(node)

    def node_message(self, node: Communication, data: message_.Message):
        message = Encoding.decode(json.dumps(data))
        self.__handle_message(message)

    def send_message(self, receiver: Node, message: str) -> None:
        self.send_to_node(receiver, message)

    def broadcast_message(self, message: str) -> None:
        self.send_to_nodes(message)

    # TODO: Implement Message subclasses (polymorphism)
    def __handle_message(self, message: message_.Message) -> None:
        if message.message_type == message_.MessageType.DISCOVERY:
            self.__handle_message_discovery(message)
        if message.message_type == message_.MessageType.TRANSACTION:
            self.__handle_message_transaction(message)
        if message.message_type == message_.MessageType.BLOCK:
            self.__handle_message_block(message)

    def __handle_message_discovery(self, message: message_.Message) -> None:
        self.peer_discovery_handler.handle_message(message)
    
    def __handle_message_transaction(self, message: message_.Message) -> None:
        self.node.handle_transaction(message.data)

    def __handle_message_block(self, message: message_.Message) -> None:
        self.node.handle_block(message.data)
