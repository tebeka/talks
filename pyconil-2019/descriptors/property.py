class Property:
    def __init__(self, func):
        self.func = func

    def __get__(self, inst, value):
        if not inst:
            return self

        return self.func(inst)


class PersonD:
    def __init__(self, first, last):
        self.first = first
        self.last = last

    @Property
    def name(self):
        return f'{self.first} {self.last}'
