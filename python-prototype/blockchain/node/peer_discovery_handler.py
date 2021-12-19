import time
from threading import Thread
from typing import Any

import blockchain.node.encoding as encoding
import blockchain.node.message as message_
import blockchain.node.socket as socket_
# import blockchain.node.socket_communication as socket_communication
from p2pnetwork.node import Node

INTERVAL_SECONDS = 10


class PeerDiscoveryHandler():

    # TODO: add typing to node: socket_communication.SocketCommunication. Fix circular import
    def __init__(self, node) -> None:
        self.communication = node

    def start(self) -> None:
        thread_status = Thread(target=self.status, args=())
        thread_discovery = Thread(target=self.discovery, args=())
        thread_status.start()
        thread_discovery.start()

    def status(self) -> None:
        while True:
            print('Current connections:')
            for peer in self.communication.peers:
                print('-', peer.host, peer.port)
            time.sleep(INTERVAL_SECONDS)

    def discovery(self) -> None:
        while True:
            handshake_message = self.handshake_message()
            self.communication.broadcast_message(handshake_message)
            time.sleep(INTERVAL_SECONDS)

    def handshake_message(self) -> str:
        socket = self.communication.socket
        peers = self.communication.peers
        msg_type = message_.MessageType.DISCOVERY
        message = message_.Message(socket, msg_type, peers)
        encoded_message = encoding.Encoding.encode(message)
        return encoded_message

    def handshake(self, node: Node) -> None:
        handshake_message = self.handshake_message()
        self.communication.send_message(node, handshake_message)

    def add_socket_to_peers(self, socket: socket_.Socket) -> None:
        if socket not in self.communication.peers:
            self.communication.peers.append(socket)

    def connect_with_node_if_not_in_peers(self, socket: socket_.Socket) -> None:
        if socket not in self.communication.peers and socket != self.communication.socket:
            self.communication.connect_with_node(socket.host, socket.port)

    def handle_message(self, message: message_.Message) -> None:
        socket = message.socket
        peers_of_peer = message.data
        self.add_socket_to_peers(socket)
        for peer_socket in peers_of_peer:
            self.connect_with_node_if_not_in_peers(peer_socket)
