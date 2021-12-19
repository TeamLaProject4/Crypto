from typing import Dict, Tuple

from blockchain.node.encoding import Encoding
from flask import Flask, jsonify, request
from flask_classful import FlaskView, route

INFO_STRING = 'Rest-API for blockchain'
NODE = None


# TODO: if python prototype continued: use FastAPI instead of Flask (https://fastapi.tiangolo.com/)
class RestAPI(FlaskView):

    def __init__(self) -> None:
        self.app = Flask(__name__)

    def start(self, port: int) -> None:
        RestAPI.register(self.app, route_base='/')
        self.app.run(host='localhost', port=port)

    # TODO: type hinting missing (circular import)
    def inject_node(self, node) -> None:
        global NODE
        NODE = node

    @route('/info', methods=['GET'])
    def info(self) -> Tuple[str, int]:
        return INFO_STRING, 200

    @route('/blockchain', methods=['GET'])
    def blockchain(self) -> Tuple[Dict, int]:
        return NODE.blockchain.to_json(), 200

    @route('memory-pool', methods=['GET'])
    def memory_pool(self) -> Tuple:
        transactions = {idx: tx.to_json() for idx, tx in enumerate(
            NODE.memory_pool.transactions)}
        return jsonify(transactions), 200

    @route('transaction', methods=['POST'])
    def transaction(self) -> Tuple:
        values = request.get_json()
        if not 'transaction' in values:
            return 'Missing transaction value', 400
        transaction = Encoding.decode(values['transaction'])
        NODE.handle_transaction(transaction)
        response = {'message': 'Transaction received'}
        return jsonify(response), 201
