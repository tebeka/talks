import sqlite3

class Proxy:
    def __init__(self, obj):
        self._obj = obj

    def __getattr__(self, attr):
        value = getattr(self._obj, attr)
        print(f'{attr} -> {value!r}')
        return value


conn = sqlite3.connect(':memory:')
proxy = Proxy(conn)
n = proxy.total_changes
print(n)

proxy.close()
