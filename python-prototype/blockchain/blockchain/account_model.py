
class AccountModel():

    def __init__(self) -> None:
        self.balances = {}

    def is_account_in_balances(self, public_key_string: str) -> bool:
        return public_key_string in self.balances.keys()

    def add_account(self, public_key_string: str):
        if not self.is_account_in_balances(public_key_string):
            self.balances[public_key_string] = 0

    def get_balance(self, public_key_string: str):
        # add account if not present in balances
        self.add_account(public_key_string)
        return self.balances[public_key_string]

    def update_balance(self, public_key_string: str, amount: int):
        # add account if not present in balances
        self.add_account(public_key_string)
        self.balances[public_key_string] += amount
