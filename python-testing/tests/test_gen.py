from operator import add

def check_add(x, y):  # <1>
    assert add(x, y) == x + y

def test_add():
    for i in range(3):
        yield check_add, i, i + 1

