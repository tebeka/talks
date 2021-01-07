from random import randint


def one_in(n):
    return randint(1, n) == 1


num_sick, num_diagnosed = 0, 0
for _ in range(1_000_000):
    is_sick = one_in(1000)
    if is_sick:
        num_sick += 1
        num_diagnosed += 1
    else:
        if one_in(20):
            num_diagnosed += 1

frac = num_sick / num_diagnosed
print(f'{frac:.2f}')
