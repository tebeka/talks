from timeit import timeit

from lex import tokenize  # noqa

text = 'If the implementation is easy to explain, it may be a good idea.'
n = 100_000

time = timeit(
    f'tokenize({text!r})',
    'from lex import tokenize',
    number=n)
per_call = time / n * 1000000
print(f'{per_call:.3f}us')
