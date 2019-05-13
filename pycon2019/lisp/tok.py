from io import BytesIO
from tokenize import tokenize
from token import tok_name

code = b'x = fib(10)'

print(f'code: {code.decode()!r}')
for tok in tokenize(BytesIO(code).readline):
    name = tok_name[tok.exact_type]
    typ = f'{tok.type:2d}/{tok_name[tok.type]}'
    etyp = f'{tok.exact_type:2d}/{tok_name[tok.exact_type]}'
    print(f'{tok.string:<5s} {typ:<15s} {etyp}')
