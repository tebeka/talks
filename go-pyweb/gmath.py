from os import path
import ctypes

here = path.dirname(path.realpath(__file__))
dll_file = f'{here}/gmath.so'


lib = ctypes.CDLL(dll_file)
lib.add.argtypes = [ctypes.c_int, ctypes.c_int]
lib.add.restype = ctypes.c_int


def add(x, y):
    if not (isinstance(x, int) and isinstance(y, int)):
        raise TypeError('parameters must be int')
    return lib.add(x, y)


def sub(x, y):
    if not (isinstance(x, int) and isinstance(y, int)):
        raise TypeError('parameters must be int')
    return lib.sub(x, y)


if __name__ == '__main__':
    print(add(1, 7))
    print(sub(1, 7))
