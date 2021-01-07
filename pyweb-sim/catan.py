from random import randint
from collections import Counter


def dice_roll():
    return randint(1, 6)


num_samples = 1_000_000
counts = Counter()
for _ in range(num_samples):
    val = dice_roll() + dice_roll()
    counts[val] += 1

for n in range(2, 13):
    frac = counts[n] / num_samples
    print(f'{n:2}: {frac:.2f}')
