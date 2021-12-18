import sys

from blockchain.node.node import Node
from blockchain.wallet import Wallet

if __name__ == '__main__':

    host = sys.argv[1]
    port = int(sys.argv[2])

    wallet = Wallet()

    node = Node(host, port, wallet)
    node.start()
