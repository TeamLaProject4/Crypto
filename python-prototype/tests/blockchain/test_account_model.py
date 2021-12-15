import pytest
from blockchain import AccountModel


def test_when_account_added_then_account_in_balances():
    public_key_string = 'alice'
    account_model = AccountModel()
    account_model.add_account(public_key_string)
    
    assert account_model.is_account_in_balances(public_key_string)


def test_when_account_added_then_account_starts_with_balance_of_zero():
    public_key_string = 'alice'
    account_model = AccountModel()
    account_model.add_account(public_key_string)
    
    assert account_model.get_balance(public_key_string) == 0
   

def test_when_balance_updated_then_balance_has_correct_amount():
    account_model = AccountModel()

    account_model.update_balance('alice', 20)
    account_model.update_balance('alice', -5)

    assert account_model.get_balance('alice') == 15
