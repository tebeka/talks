import fields


class Trade:
    symbol = fields.Symbol()
    price = fields.Price()
    volume = fields.Volume()

    def __init__(self, symbol, price, volume):
        self.symbol = symbol
        self.price = price
        self.volume = volume

    def __repr__(self):
        cls = self.__class__.__name__
        return f'{cls}({self.symbol!r}, {self.price!r}, {self.volume!r})'
