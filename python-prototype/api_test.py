from blockchain import TxType
from blockchain.wallet import Wallet
from blockchain.node.encoding import Encoding
import requests

if __name__ == '__main__':
    alice = Wallet()
    bob = Wallet()
    exchange = Wallet()

    exchange_tx = exchange.create_transaction(
        alice.public_key_string(), 50, TxType.EXCHANGE)
    
    url = 'http://localhost:5000/transaction'
    body = {'transaction': Encoding.encode(exchange_tx)}
    request = requests.post(url, json=body)
    print(request.text)
