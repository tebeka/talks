callbacks = [lambda i=i: i for i in range(5)]  # <1>
print([c() for c in callbacks])
# => [0, 1, 2, 3, 4]


def make_callback(i):  # <2>
    return lambda: i


callbacks = [make_callback(i) for i in range(10)]
print([c() for c in callbacks])
# => [0, 1, 2, 3, 4]
