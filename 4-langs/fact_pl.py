def fact(n):
    if n == 1:
        return n
    n1 = n - 1
    return n * fact(n1)


print(fact(10))
