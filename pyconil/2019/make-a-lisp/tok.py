from io import BytesIO
from tokenize import tokenize
from token import tok_name


def print_tokens(code):
    tokens = tokenize(BytesIO(code).readline)
    for tok in tokens:
        name = tok_name[tok.exact_type]
        print(f'{tok.string:<10s} {name}')
