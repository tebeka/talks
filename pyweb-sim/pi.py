"""Calculate π using simulation"""
from random import uniform
from math import sqrt

R = 1

n, inner = 100_000_000, 0
for _ in range(n):
    x, y = uniform(0, 1), uniform(0, 1)
    # Point falls inside the circle
    if sqrt(x*x + y*y) < R:
        inner += 1

ratio = inner / n
# Since R = 1, the ratio of inner points from all the points is 1/4 of a
# circle
print(4 * ratio)
