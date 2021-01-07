from random import uniform
from math import sqrt

n, inner = 100_000_000, 0
for _ in range(n):
    x, y = uniform(0, 1), uniform(0, 1)
    if sqrt(x*x + y*y) < 1:
        inner += 1
ratio = inner/n
print(4*ratio)
