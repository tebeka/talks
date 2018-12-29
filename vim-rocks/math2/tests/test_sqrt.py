import pytest
from math2.sqrt import sqrt
import json
from os import path

here = path.dirname(path.abspath(__file__))


def load_cases():
    with open(f'{here}/sqrt_cases.json') as fp:
        for case in json.load(fp):
            yield case['value'], case['result']


@pytest.mark.parametrize('val, result', load_cases())
def test_sqrt(val, result):
    assert round(sqrt(val), 4) == round(result, 4), 'bad sqrt'
