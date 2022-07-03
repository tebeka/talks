import math

from calc import calc


def test_calc():
    val = calc('2 * pi')
    assert math.tau == val
