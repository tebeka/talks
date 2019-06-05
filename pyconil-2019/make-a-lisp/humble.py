import operator


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


builtin = {
    '+': operator.add,
    '-': operator.sub,
    '*': operator.mul,
    '/': operator.truediv,
    '%': operator.mod,
    '>': operator.gt,
    '<': operator.lt,
    '=': operator.eq,
}


def evaluate(expr):
    # number: 2.718
    if isinstance(expr, float):
        return expr

    # name: a, * ...
    if isinstance(expr, str):
        return builtin[expr]

    # If we're here, it a list
    op, *rest = expr

    # (and (> 1 2) (> 3 4))
    if op == 'and':
        for expr in rest:
            if not evaluate(expr):
                return 0.0
        return 1.0

    # (or (< 7 3) (> 3 2) 17)
    if op == 'or':
        for expr in rest:
            if evaluate(expr):
                return 1.0
        return 0.0

    # (if (> 1 2) 10 20)
    if op == 'if':
        # TODO: Support if without else
        cond, true_expr, false_expr = rest
        expr = true_expr if evaluate(cond) else false_expr
        return evaluate(expr)

    # (* 3 7)
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
            s_expr = reader(code)
            val = evaluate(s_expr)
            print(val)
        except (EOFError, KeyboardInterrupt):
            break
        except Exception as err:
            print(f'ERROR: {err.__class__.__name__} {err}')
    print('Yalla Bye ☺')


if __name__ == '__main__':
    import readline  # History with arrows, paren highlight ...# noqa
    repl()

# code = '(if (even? n) (/ n 2) (+ (* 3 n) 1))'
# print(reader(code))
