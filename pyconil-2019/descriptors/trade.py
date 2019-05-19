import fields


class Trade:
    symbol = fields.Symbol()
    price = fields.Price()
    volume = fields.Volume()

    def __init__(self, symbol, price, volume):
        self.symbol = symbol
        self.price = price
        self.volume = volume
