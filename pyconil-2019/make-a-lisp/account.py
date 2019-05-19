def make_account(balance):
    # MT: Python 2 hack with b = [balance] (account2.py)
    def account(amount):
        nonlocal balance

        balance += amount
        return balance

    return account


acct = make_account(100)
