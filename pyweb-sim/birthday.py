"""Simulation of the "Birday problem"

See https://en.wikipedia.org/wiki/Birthday_problem for a description of the
problem.
"""
from random import randint


def random_birthday():
    """Return a birthday as day-of-year."""
    return randint(1, 365)


def dup_in_group(size):
    """Return True if there's at least one duplicate birthday in a random
    group of size.
    """
    group = set()
    for _ in range(size):
        day = random_birthday()
        if day in group:
            return True
        group.add(day)
    return False


n, group_size, dups = 1_000_000, 23, 0
for _ in range(n):
    if dup_in_group(group_size):
        dups += 1

frac = dups / n
print(f'{frac:.2f}')
