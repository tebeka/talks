from os import path
import ctypes

here = path.dirname(path.realpath(__file__))
dll_file = f'{here}/fetch.so'


lib = ctypes.CDLL(dll_file)
lib.fetch.argtypes = [ctypes.c_char_p]
lib.fetch.restype = ctypes.c_char_p


def fetch(url):
    out = lib.fetch(url.encode())
    if out is None:
        raise IOError(url)
    html = out.decode()
    return html

if __name__ == '__main__':
    html = fetch('http://golang.org')
    print(html)
