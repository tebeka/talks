callbacks = [lambda: i for i in range(5)]
print([c() for c in callbacks])
# => [4, 4, 4, 4, 4]
