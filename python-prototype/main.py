import sys

from blockchain.node.node import Node
from blockchain.wallet import Wallet

if __name__ == '__main__':

    host = sys.argv[1]
    p2p_port = int(sys.argv[2])
    rest_api_port = int(sys.argv[3])

    wallet = Wallet()

    node = Node(host, p2p_port, wallet)
    node.start_p2p()
    node.start_rest_api(rest_api_port)
