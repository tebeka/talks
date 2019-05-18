from dis import dis


def collatz(n):
    """Return the next number of Collatz conjucture"""
    if n % 2 == 0:
        return n / 2
    return n * 3 + 1
