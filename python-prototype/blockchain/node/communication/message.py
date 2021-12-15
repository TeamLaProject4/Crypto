

class Message():

    def __init__(self, socket, msg_type, data) -> None:
        self.sender_connector = socket
        self.msg_type = msg_type
        self.data = data
