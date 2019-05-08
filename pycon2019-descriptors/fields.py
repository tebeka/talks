class Field:
    def __get__(self, inst, owner):
        if inst is None:
            return self
        return getattr(inst, self._attr)

    def __set__(self, inst, value):
        self.assert_valid(value)
        setattr(inst, self._attr, value)

    def __set_name__(self, owner, name):
        self._attr = f'_{owner.__name__}_{name}'

    def assert_valid(self, value):
        pass


class Symbol(Field):
    def assert_valid(self, value):
        if not str.isupper(value):
            raise ValueError(f'symbol not upper case: {value!r}')


class Price(Field):
    def assert_valid(self, value):
        if not isinstance(value, float):
            raise TypeError(f'price not a float: {value!r}')

        if value <= 0:
            raise ValueError(f'negative price: {value!r}')


class Volume(Field):
    def assert_valid(self, value):
        if not isinstance(value, int):
            raise TypeError(f'volume not an int: {value!r}')

        if value <= 0:
            raise ValueError(f'negative volume: {value!r}')
