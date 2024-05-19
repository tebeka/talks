import json


class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    @classmethod
    def from_json(cls, data):
        obj = json.loads(data)
        return cls(obj['x'], obj['y'])
