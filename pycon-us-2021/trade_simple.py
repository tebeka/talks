class Trade:
    def __init__(self, symbol, volume, price, buy):
        self.symbol = symbol
        self.volume = volume
        self.price = price
        self.buy = buy


tr1 = Trade('BRK-A', 1, 321801.07, True)
print(f'{tr1!r}')

# Trade(symbol='BRK-A', volume=1, price=321801.07, buy=True)
