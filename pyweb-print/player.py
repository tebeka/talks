import re


class Player:
    def __init__(self, name, mana, keys):
        self.name = name
        self.mana = mana
        self.keys = keys

    def __repr__(self):
        # __repr__ is for developers, usually *a* way to create such an object
        # self can be instance of this class or any subclass of it
        cls = self.__class__.__name__
        return f'{cls}({self.name!r}, {self.mana!r}, {self.keys!r})'

    def __str__(self):
        # __str__ is for users
        return f'{self.name} has {self.keys}'

    def __format__(self, spec):
        return re.sub(r'%[a-z]', self._sub, spec)

    def _sub(self, match):
        name = match.group()[1:]  # remove %
        attr = {
            'n': 'name',
            'm': 'mana',
            'k': 'keys',
        }.get(name, '')
        return str(getattr(self, attr, match.group()))


p1 = Player('Parzival', 392, {'copper', 'jade'})
