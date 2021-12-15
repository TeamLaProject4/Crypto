from typing import Any

import jsonpickle


class Encoding():

    @staticmethod
    def encode(data: Any) -> str:
        return jsonpickle.encode(data, unpicklable=True)

    @staticmethod
    def decode(data: Any) -> Any:
        return jsonpickle.decode(data)
