import json
from typing import Any

from Crypto.Hash import SHA256


class Utils():

    @staticmethod
    def hash(data: Any) -> SHA256.SHA256Hash:
        data_string = json.dumps(data)
        data_bytes = data_string.encode('utf-8')
        data_hash = SHA256.new(data_bytes)
        return data_hash
