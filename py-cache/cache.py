from functools import wraps


def cached(fn):
    cache = {}  # args -> value

    @wraps(fn)
    def wrapper(*args, **kw):
        if args in cache:
            return cache[args]

        value = cache[args] = fn(*args)
        return value

    return wrapper
