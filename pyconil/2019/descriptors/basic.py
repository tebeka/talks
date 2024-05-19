class Desc:
    def __get__(self, inst, owner):
        print('__get__: inst=%r, owner=%r' % (inst, owner))

    def __set__(self, inst, value):
        print('__set__: inst=%r, value=%r' % (inst, value))


class Stock:
    symbol = Desc()
