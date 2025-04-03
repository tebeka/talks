# So You Think You Can Print?

## Ideas
- str vs repr
- f-string
    - ^<>
    - .2f
    - 08b, e,f,g  (mention 0x, 0b, 0o, `_`)
    - n = 0.7, f'{n:%}'
    - f'{name:>{n}} `os.get_termina_size`
    - f'{n:,}
    - https://docs.python.org/3/library/string.html#formatspec
    - print('This is \N{radioactive sign}')
    - f'{now:%Y%m%d}'
- print
    - end=''
    - file=
    - flush=True (example with subprocess)
    - print('a', 'b', 'c', sep=',')
- '\n'.join(generator) for templating
- `__str__`, `__repr__`, `__format__`
    - Prevent printing sensitive data
- logging %s (performance)
    - https://docs.python.org/3/howto/logging-cookbook.html#use-of-alternative-formatting-styles
    - print will slow you down
	- time python spammer.py 10000
	- time python spammer.py 10000 > /tmp/spam.log
- Unicode?
- Python 3 made print function
- print only in main, otherwise log
