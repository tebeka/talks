from dataclasses import dataclass


@dataclass
class Person:
    first: str
    last: str

    @property
    def name(self):
        return f'{self.first} {self.last}'
