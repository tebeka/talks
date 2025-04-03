from functools import wraps


def metrics(fn):
    call_count = 0

    @wraps(fn)
    def wrapper(*args, **kw):
        call_count += 1
        print(f'{fn.__name__} called {call_count} times')

        return fn(*args, **kw)

    return wrapper


@metrics
def handler(message):
    print(f'got {message}')
