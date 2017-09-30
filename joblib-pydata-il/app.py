from joblib import Memory
from os.path import expanduser

memory = Memory(cachedir=expanduser('~/.cache/myapp'), verbose=0)


@memory.cache
def fib(n):
    print(f'fib({n})')
    if n < 2:
        return 1
    return fib(n-1) + fib(n-2)

print(fib(10))
