from dataclasses import dataclass


class Property:
    def __init__(self, func):
        self.func = func

    def __get__(self, inst, value):
        if not inst:
            return self

        return self.func(inst)


@dataclass
class Person:
    first: str
    last: str

    @Property
    def name(self):
        return f'{self.first} {self.last}'
