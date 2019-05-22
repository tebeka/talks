from abc import ABC, abstractmethod


class Field(ABC):
    def __get__(self, inst, owner):
        if inst is None:
            return self
        return getattr(inst, self._attr)

    def __set__(self, inst, value):
        self.assert_valid(value)
        setattr(inst, self._attr, value)

    # x = Field() -> Field.__set_name__(Field, 'x')
    def __set_name__(self, owner, name):
        self._attr = f'_{owner.__name__}_{name}'

    @abstractmethod
    def assert_valid(self, value):
        ...
