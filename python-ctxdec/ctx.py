## Greeter
class greeter(object):
    def __enter__(self):
        print('Hai')

    def __exit__(self, exc_type, exc_value, traceback):
        print('Bai')


## timed_block
from time import time


class timed_block(object):
    def __init__(self, name):
        self.name = name

    def __enter__(self):
        self.start = time()

    def __exit__(self, exc_type, exc_value, traceback):
        duration = time() - self.start
        print('{} took {:0.2f}sec'.format(self.name, duration))


## Closer
class closing(object):
    def __init__(self, obj):
        self.obj = obj

    def __enter__(self):
        pass

    def __exit__(self, exc_type, exc_value, traceback):
        self.obj.close()


## Database
class dbctx(object):
    def __init__(self, conn):
        self.conn = conn

    def __enter__(self):
        return self.conn.cursor()

    def __exit__(self, exc_type, exc_value, traceback):
        if exc_type is None:
            print('Committing')
            self.conn.commit()
        else:
            print('Rolling back')
            self.conn.rollback()


## contextlib
from contextlib import contextmanager


@contextmanager
def dbctx(conn):
    try:
        yield conn.cursor()
        conn.commit()
    except:
        conn.rollback()
        raise  # Need to re-raise
