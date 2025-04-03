"""Simulate percision of a medical test.

For more details see
https://psychscenehub.com/psychinsights/well-understand-probabilities-medicine/
"""
from random import randint


def one_in(n):
    """Return True on in n cases."""
    return randint(1, n) == 1


def diagnose(is_sick):
    """Return result is diagnostic."""
    if is_sick:
        # We're 100% correct in sick people.
        return True

    # The test of a disease presents a rate of 5% (1 in 20) false positives.
    # (false positive = healthy diagnosed as sick)
    return one_in(20)


num_sick, num_diagnosed = 0, 0
for _ in range(1_000_000):
    # The disease strikes 1/1000 of the population.
    is_sick = one_in(1000)
    if is_sick:
        num_sick += 1

    if diagnose(is_sick):
        num_diagnosed += 1

frac = num_sick / num_diagnosed
print(f'{frac:.2f}')
