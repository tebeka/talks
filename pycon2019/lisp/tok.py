from io import BytesIO
from tokenize import tokenize
from token import tok_name


def show_tokens(code):
    print(f'code: {code.decode()!r}')
    for tok in tokenize(BytesIO(code).readline):
        name = tok_name[tok.exact_type]
        print(f'{tok.string:<5s} {name}')
