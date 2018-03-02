from operator import mul
from functools import reduce

def fact(n):
    return reduce(mul, range(1, n+1))


print(fact(10))
