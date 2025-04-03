from functools import wraps


def log_errors(fn):
    @wraps(fn)
    def wrapper(*args, **kw):
        try:
            val = fn(*args, **kw)
        except Exception as err:
            pass

        print(f'{fn.__name__}: error={err}')
        if err:
            raise err
        return val

    return wrapper


@log_errors
def div(a, b):
    return a / b
