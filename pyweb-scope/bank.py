def account(balance):
    def deposit(amount):
        nonlocal balance

        new = balance + amount
        print(f'{balance} -> {new}')
        balance = new

    return deposit
