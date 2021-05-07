from dataclasses import dataclass


@dataclass
class Trade:
    symbol: str
    volume: int
    price: float
    buy: bool


tr1 = Trade('BRK-A', 1, 321801.07, True)
print(f'{tr1!r}')

# Trade(symbol='BRK-A', volume=1, price=321801.07, buy=True)
