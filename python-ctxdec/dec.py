## Greeter
def greeter(fn):
    def wrapper():
        print('Hello {}'.format(fn.__name__))
        fn()
        print('Bye {}'.format(fn.__name__))

# TypeError
# need to return wrapper

# TypeError
# x, y (not generic, move to *args, **kw)

# See None
# Need to return value
# val = fn(x, y)
# try/finally

# help(add)
# from functools import wraps

from functools import wraps


def greeter(fn):
    @wraps(fn)
    def wrapper(*args, **kw):
        print('Hello {}'.format(fn.__name__))
        try:
            return fn(*args, **kw)
        finally:
            print('Bye {}'.format(fn.__name__))

    return wrapper

## Timing
from functools import wraps
from time import time


def timed(fn):
    @wraps(fn)
    def wrapper(*args, **kw):
        start = time()
        try:
            return fn(*args, **kw)
        finally:
            duration = time() - start
            print('{} took ({:.2f}sec)'.format(fn.__name__, duration))

    return wrapper


## Bonus
from ctx import timed_block


def timed(fn):
    @wraps(fn)
    def wrapper(*args, **kw):
        with timed_block(fn.__name__):
            return fn(*args, **kw)
    return wrapper


## Cache
def cached(fn):
    cache = {}
    not_found = object()

    def wrapper(*args, **kw):
        value = cache.get(args, not_found)
        if value is not_found:
            value = cache[args] = fn(*args, **kw)

        return value

    return wrapper


## Multiply
def mulby(n):
    def wrapper(fn):
        @wraps(fn)
        def wrapped(*args, **kw):
            return n * fn(*args, **kw)

        return wrapped
    return wrapper
