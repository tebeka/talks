class Id(object):
    _next_id = 0
    def __init__(self):
        self.id, Id._next_id = Id._next_id, Id._next_id + 1

class PointId(Id):
    def __init__(self, x, y):
        super(PointId, self).__init__()  # <1>
        self.x, self.y = x, y
    def __repr__(self):
        return 'Point#{}: {},{}'.format(self.id, self.x, self.y)


pi = PointId(1, 2)
print(pi)
# => Point#0: 1,2
