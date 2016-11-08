def sub(x, y):
    """Returns the y's predecessor to x"""
    return x - y


sub(1, 7)
sub(y=7, x=1)
sub(1, y=7)
sub(x=1, 7)  # syntax error

args = 1, 7
sub(*args)
# * in call unpacks sequence to positional

a, b = (1, ), (7, )
sub(*(a + b))
# new in 3.5
sub(*a, *b)

def vargs(*args):
    print(args)

vargs()
vargs(1, 2, 3)

# * in definition packs positional to tuple

kw = {'x': 1, 'y': 7}
sub(**kw)
# ** in call unpacks mapping to keyword arguments

k = {'x': 1}
w = {'y': 7}
kw = k.copy()
kw.update(w)
sub(**kw)

# new in 3.5, beware of doubles
sub(**k, **w)


def kwargs(**kw):
    print(kw)

kwargs()
kwargs(name='daffy', age=78)
# ** in function call packs keyword arguments to dict


def eat_all(*args, **kw):
    print(args)
    print(kw)

eat_all()
eat_all('bugs')
eat_all(taz='brown')
eat_all('daffy', 'taz', colors=['black', 'brown'])

# very bad design, mostly used in decorators


def sub(x, y=1):
    return x - y

sub(1, 7)
sub(8)

# Use var=value to specify default arguments
# NEVER USE MUTABLE DEFAULT ARGUMENTS


class subber:
    def __init__(self, n):
        self.n = n

    def __call__(self, x):
        return x - self.n

sub7 = subber(7)
sub7(1)

# any object which implements __call__ can be used as a function
sub.__call__(1, 7)


def sub(x: int, y: int) -> int:
    return x - y

sub(1, 2)
help(sub)
sub.__annotations__


def sub(x, y, *, verbose=False):
    if verbose:
        print('subing %s to %s' % (x,  y))
    return x - y

# ensure keyword only parameters (start without * and show sub(1, 2, True) not
# that readable)

sub(1, 7)
sub(1, 7, True)  # Error
sub(1, 7, verbose=True)


# PEP 492
async def sub(x, y):
    return x - y
sub(1, 7)  # return coroutine


cat = ''.join
cat(['a', 'b', 'c'])

# You can use bound methods as functions

from functools import partial

compact_ws = partial(re.compile('\s+').sub, ' ')
compact_ws('how   are  \t you?')

# You can use functool.partial to create new functions

sub = lambda x, y: x - y
sub(1, 7)

# You can use lambda to create adhoc functions. Very limited capability

fns = [lambda: n for n in range(5)]
vals = [fn() for fn in fns]
print(vals)

fns = [lambda n=n: n for n in range(5)]
vals = [fn() for fn in fns]
print(vals)

# default arguments computed at definition time

def prappend(n, vals=[]):
    vals.append(n)
    print(vals)
prappend(1)
prappend(2)

def prappend(n, vals=None):
    vals = [] if vals is None else vals
    vals.append(n)
    print(vals)
prappend(1)
prappend(2)

# never use mutable default arguments


def make_subber(n):
    def subber(x):
        return x - n
    return subber

# closure

sub7 = make_subber(7)
sub7(1)



ncalls = 0
def sub(x, y):
    ncalls += 1
    return x - y

sub(1, 7)  # UnboundLocalError


ncalls = 0
def sub(x, y):
    global ncalls
    ncalls += 1
    return x - y
sub(1, 7)
print(ncalls)


from time import time
call_times = []
def sub(x, y):
    call_times.append(time())
    return x - y
sub(1, 7)
print(call_times)


def sub(x, y):
    """Returns the y's predecessor to x

    >>> sub(10, 3)
    7
    >>> sub(5, 22)
    -17
    """
    return x - y

import doctest
doctest.testmod()

# Sphinx can run doctest - test your documentation!
# can run tests of text files
# nose & py.test can run doctest tests as well

from functools import wraps
from time import time, sleep

def timed(fn):
    @wraps(fn)
    def wrapper(*args, **kw):
        start = time()
        try:
            return fn(*args, **kw)
        finally:
            print('%s took %.2fsec' % (fn.__name__, time() - start))
    return wrapper

@timed
def sub(x, y):
    sleep(0.1)
    return x - y

sub(1, 7)

def ncalls(fn):
    count = 0
    @wraps(fn)
    def wrapper(*args, **kw):

        count += 1
        print('%s called %d times' % (fn.__name__, count))
        return fn(*args, **kw)
    return wrapper


@ncalls
def sub(x, y):
    return x - y

sub(1, 6)
# UnboundLocalError

def ncalls(fn):
    count = 0
    @wraps(fn)
    def wrapper(*args, **kw):
        nonlocal count

        count += 1
        print('%s called %d times' % (fn.__name__, count))
        return fn(*args, **kw)
    return wrapper

sub(1, 6)
sub(1, 6)

# you can use decorators to add functionality to functions

# https://www.youtube.com/watch?v=aintdHnqaio
