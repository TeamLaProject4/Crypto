from __future__ import annotations

from typing import Any

from p2pnetwork.node import Node


class SocketCommunication(Node):

    def __init__(self, host: str, port: int) -> None:
        super().__init__(host, port, callback=None)

    def start(self) -> None:
        super().start()

    def inbound_node_connected(self, node: SocketCommunication) -> None:
        print('inbound connection')
        self.send_to_node(node, 'Hi, I am the node you connected to')

    def outbound_node_connected(self, node: SocketCommunication) -> None:
        print('outbound connection')
        self.send_to_node(node, 'Hi, I am the node who initialized the connection')

    def node_message(self, node: SocketCommunication, data: Any):
        print(data)
