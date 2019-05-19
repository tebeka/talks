from io import BytesIO
from tokenize import tokenize
from token import tok_name


def print_tokens(code):
    for tok in tokenize(BytesIO(code).readline):
        name = tok_name[tok.exact_type]
        print(f'{tok.string:<10s} {name}')
