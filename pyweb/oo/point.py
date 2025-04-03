from textwrap import dedent

class Point:
    pass


code = '''
    print('making a Point')
    def __init__(self, x, y):
        self.x, self.y = x, y

    def move(self, dx, dy):
        self.x += dx
        self.y += dy
'''

env = {}
exec(dedent(code), env)
for k, v in env.items():
    setattr(Point, k, v)


p = Point(1, 2)
p.move(2, 4)
print(p.x, p.y)
