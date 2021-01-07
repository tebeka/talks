from random import randint


def random_birthday():
    return randint(1, 365)


def dup_in_group(size):
    seen = set()
    for _ in range(size):
        day = random_birthday()
        if day in seen:
            return True
        seen.add(day)
    return False


n, group_size, dups = 1_000_000, 23, 0

for _ in range(n):
    if dup_in_group(group_size):
        dups += 1

frac = dups / n
print(f'{frac:.2f}')
