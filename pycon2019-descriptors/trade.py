from dataclasses import dataclass

import fields


@dataclass
class Trade:
    symbol: fields.Field = fields.Symbol()
    price: fields.Field = fields.Price()
    volume: fields.Field = fields.Volume()

    def __repr__(self):
        cls = self.__class__.__name__
        return f'{cls}({self.symbol!r}, {self.price!r}, {self.volume!r})'
