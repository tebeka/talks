from sys import stderr


def setup():
    stderr.write('Let the games begin!\n')  # <1>


def teardown():
    stderr.write('\nAre we there yet?\n')


def test_ok():
    assert 1 == 1, 'okey dokey'


def test_bad():
    assert 1 == 2, 'please reinstall universe and reboot'
