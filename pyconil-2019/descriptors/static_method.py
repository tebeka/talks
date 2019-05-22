class StaticMethod:
    def __init__(self, func):
        self.func = func

    def __get__(self, inst, owner):
        return self.func


class MathD:
    @StaticMethod
    def neg(val):
        return -val
