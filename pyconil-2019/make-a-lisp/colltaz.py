def collatz(n):
    """Return the next number in Collatz conjucture"""
    if n % 2 == 0:
        return n / 2
    return n * 3 + 1
