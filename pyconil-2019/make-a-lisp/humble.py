def tokenize(code):
    code = code.replace('(', ' ( ')
    code = code.replace(')', ' ) ')
    return code.split()


def read_sexpr(tokens):
    if not tokens:
        raise EOFError

    tok = tokens.pop(0)
    # tok, tokens = tokens[0], tokens[1:]
    if tok == '(':
        child = []
        while tokens[0] != ')':
            child.append(read_sexpr(tokens))
        tokens.pop(0)  # Remove closing )
        return child

    if tok == ')':
        raise SyntaxError('unexpected )')

    # If we're here, it's an atom (name or number)
    try:
        return float(tok)
    except ValueError:
        return tok  # name


def reader(code):
    return read_sexpr(tokenize(code))

import operator
builtin = {
    '*': operator.mul,
}

# TODO: Add /, %, +, -
# Back 3:20


def evaluate(expr):
    # number: 2.718
    if isinstance(expr, float):
        return expr

    # name: a, * ...
    if isinstance(expr, str):
        return builtin[expr]

    # (* 3 7)
    if isinstance(expr, list):
        op, *rest = expr
        # Python 2
        # op, rest = expr[0], expr[1:]
        func = evaluate(op)
        args = [evaluate(arg) for arg in rest]
        return func(*args)





# Read, Eval, Print, Loop
def repl():
    while True:
        try:
            code = input('» ')
            if not code.strip():
                continue
            # TODO: tokenize, read s-expression, print tree
            s_expr = reader(code)
            val = evaluate(s_expr)
            print(val)
        except (EOFError, KeyboardInterrupt):
            break
    print('Yalla Bye ☺')


if __name__ == '__main__':
    import readline  # History with arrows, paren highlight ...# noqa
    repl()

# code = '(if (even? n) (/ n 2) (+ (* 3 n) 1))'
# print(reader(code))
