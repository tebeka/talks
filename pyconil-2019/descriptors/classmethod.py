import json
from functools import partial


class ClassMethod:
    def __init__(self, func):
        self.func = func

    def __get__(self, inst, owner):
        return partial(self.func, owner)


class PointD:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    @ClassMethod
    def from_json(cls, data):
        obj = json.loads(data)
        return cls(obj['x'], obj['y'])
