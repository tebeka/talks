def test_range():
    items = range(10)
    size = len(items)
    expected = 11

    assert size == expected, 'bad range'
