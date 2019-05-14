from io import BytesIO
from tokenize import tokenize
from token import tok_name

code = b'x = collatz(7)'

print(f'code: {code.decode()!r}')
for tok in tokenize(BytesIO(code).readline):
    name = tok_name[tok.exact_type]
    print(f'{tok.string:<5s} {name}')
