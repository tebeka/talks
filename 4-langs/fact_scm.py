def fact(n):

    def _fact(n, acc):
        if n == 1:
            return acc
        return _fact(n-1, acc*n)

    return _fact(n, 1)

print(fact(10))
