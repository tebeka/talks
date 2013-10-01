import pickle

class Point(object):  # <1>
    def __init__(self, x, y):
        self.x, self.y = x, y

p1, p2 = Point(1, 1), Point(2, 2)

tmp = '/tmp/points.p'
with open(tmp, 'wb') as out:
    pickle.dump(p1, out)
    pickle.dump(p2, out)

with open(tmp) as fo:
    p1a, p2a = pickle.load(fo), pickle.load(fo)  # <2>

assert p2a.y == p2.y, 'corrupted data'
