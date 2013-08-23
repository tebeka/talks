class Id(object):
    _next_id = 0
    def __init__(self):
        self.id, Id._next_id = Id._next_id, Id._next_id + 1

class PointId(Id):
    def __init__(self, x, y):
        self.x, self.y = x, y
    def __repr__(self):
        return 'Point#{}: {},{}'.format(self.id, self.x, self.y)

pi = PointId(1, 2)
print(pi)
# Traceback (most recent call last):
#   File "src/ctor.py", line 19, in <module>
#     print(pi)
#   File "src/ctor.py", line 15, in __repr__
#     return 'Point#{}: {},{}'.format(self.id, self.x, self.y)
# AttributeError: 'PointId' object has no attribute 'id'

