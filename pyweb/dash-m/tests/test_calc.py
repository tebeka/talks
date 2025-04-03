import math

from pytest import approx

from calc import calc


def test_calc():
    val = calc('2 * pi')
    assert val == approx(math.tau)
