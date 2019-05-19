from field import Field


class Symbol(Field):
    def assert_valid(self, value):
        if not value:
            raise ValueError('empty symbol')

        if not str.isupper(value):
            raise ValueError(f'symbol not upper case: {value!r}')


class Price(Field):
    def assert_valid(self, value):
        if not isinstance(value, float):
            raise TypeError(f'price not a float: {value!r}')

        if value <= 0:
            raise ValueError(f'non positive price: {value!r}')


class Volume(Field):
    def assert_valid(self, value):
        if not isinstance(value, int):
            raise TypeError(f'volume not an int: {value!r}')

        if value <= 0:
            raise ValueError(f'non positive volume: {value!r}')
