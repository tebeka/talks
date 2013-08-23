class AE(object):
    def __init__(self, i):
        self.i = i

    def __eq__(self, other):  # <1>
        return True
