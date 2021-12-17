from __future__ import annotations

from typing import Any, List

from p2pnetwork.node import Node

import blockchain.node.communication.socket as socket
import blockchain.node.communication.peer_discovery_handler as peer_discovery

# TODO: refactor bootnode. Bootnodes should be in a config file or so and must be derived to a list
BOOTNODE = socket.Socket('localhost', 5000)


class SocketCommunication(Node):

    # TODO: remove hardcoded BOOTNODE
    def __init__(self, host: str, port: int, bootnodes: List[socket.Socket] = [BOOTNODE]) -> None:
        super().__init__(host, port, callback=None)
        self.peers = []
        self.peer_discovery_handler = peer_discovery.PeerDiscoveryHandler(self)
        self.socket = socket.Socket(host, port)
        self.bootnodes = bootnodes

    def start(self) -> None:
        super().start()
        self.peer_discovery_handler.start()
        self.connect_with_bootnode()
    
    def connect_with_bootnode(self) -> None:
        # TODO: remove hardcoded indice 0, instead try all bootnodes, one by one till connected
        bootnode = self.bootnodes[0]
        if self.socket != bootnode:
            self.connect_with_node(bootnode.host, bootnode.port)

    def inbound_node_connected(self, node: SocketCommunication) -> None:
        self.peer_discovery_handler.handshake(node)

    def outbound_node_connected(self, node: SocketCommunication) -> None:
        self.peer_discovery_handler.handshake(node)

    def inbound_message(self, node: SocketCommunication, data: Any):
        print(data)

    def send_message(self, receiver: Node, message: str) -> None:
        self.send_to_node(receiver, message)

    def broadcast_message(self, message: str) -> None:
        self.send_to_nodes(message)
