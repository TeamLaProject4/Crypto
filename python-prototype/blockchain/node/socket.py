from __future__ import annotations


class Socket():

    def __init__(self, host: str, port: int) -> None:
        self.host = host
        self.port = port

    # TODO rewrite method as: (self, __o: object)
    def __eq__(self, socket: Socket) -> bool:
        return (self.host == socket.host) and (self.port == socket.port)
