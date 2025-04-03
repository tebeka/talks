#!/usr/bin/env python
"""Parallel map"""

# The big advantage of this implementation is that "fork" is very fast on
# copying data, so if you pass big arrays as arguments and return small values
# this is a win

# FIXME
# * For too many items, we get "Too many open files"
# * Handle child exceptions

import pickle
import os


def spawn(func, data):
    """Spawn a child process which run func on data"""
    rdr, wtr = os.pipe()
    rdr, wtr = os.fdopen(rdr, 'rb'), os.fdopen(wtr, 'wb')
    pid = os.fork()
    if pid:  # parent
        return pid, rdr

    # child
    pickle.dump(func(data), wtr)
    wtr.close()
    raise SystemExit


def wait(child):
    """Wait for child to finish, load data from pipe"""
    pid, fp = child
    with fp:
        os.waitpid(pid, os.P_WAIT)
        return pickle.load(fp)


def pmap(func, items):
    """Parallel map (using fork) of func over items

    >>> pmap(lambda x: x * 2, range(5))
    [0, 2, 4, 6, 8]
    """
    children = [spawn(func, item) for item in items]
    return [wait(child) for child in children]


if __name__ == '__main__':
    def fib(n):
        a, b = 1, 1
        while n > 1:
            a, b = b, a + b
            n -= 1
        return b

    nums = range(10)
    print('pmap(fib, %s)' % str(nums))
    for num, val in zip(nums, pmap(fib, nums)):
        print('%d -> %d' % (num, val))
