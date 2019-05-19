def make_account(balance):
    balance = [balance]

    def account(amount):
        balance[0] += amount
        return balance[0]

    return account


acct = make_account(100)
