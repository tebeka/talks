from dis import dis


def collatz(n):
    if n % 2 == 0:
        return n / 2
    return n * 3 + 1


print(dis(collatz))
