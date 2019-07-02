from dataclasses import dataclass


@dataclass
class Trade:
    symbol: str
    volume: int
    price: float
    buy: bool


t = Trade('BRK-A', 1, 321801.07, True)
print(f'{t!r}')

# Trade(symbol='BRK-A', volume=1, price=321801.07, buy=True)
