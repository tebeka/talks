def logger(fn):
    def wrapper(*args, **kw):
        print('{} called'.format(fn.__name__))
        return fn(*args, **kw)

    return wrapper

@logger
def add(x, y):
    '''Adds x to y'''
    return x + y

add(1, 2)
# => add called
print(add.__doc__)
# => None
