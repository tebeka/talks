"""Full humble lisp implementation"""
import re
import operator
from collections import ChainMap
from dataclasses import dataclass


def tokenize(code):
    code = re.sub(r';.*', '', code)  # nuke comments
    for c in '()':
        code = code.replace(c, f' {c} ')
    return code.split()


def read_sexpr(tokens):
    if not tokens:
        raise EOFError

    tok = tokens.pop(0)
    if tok == '(':
        child = []
        while tokens[0] != ')':
            child.append(read_sexpr(tokens))
        tokens.pop(0)  # remove closing ')'
        return child

    # TODO: file:line
    if tok == ')':
        raise SyntaxError('unexpected )')

    if tok in {'#t', '#f'}:
        return tok == '#t'

    try:
        return float(tok)
    except ValueError:
        return tok


def reader(code):
    # TODO: Multiple s-expressions
    tokens = tokenize(code)
    return read_sexpr(tokens)


@dataclass
class Lambda:
    args: list
    body: list
    env: ChainMap  # Closure

    def __call__(self, *params):
        args = {name: val for name, val in zip(self.args, params)}
        env = ChainMap(args, self.env)
        return evaluate(self.body, env)

    def __repr__(self):
        args = ' '.join(self.args)
        body = lispify(self.body)
        return f'(lambda ({args}) {body})'


def begin(*args):
    if args:
        return args[-1]


builtin = ChainMap({
    '+': operator.add,
    '-': operator.sub,
    '*': operator.mul,
    '/': operator.truediv,
    '>': operator.gt,
    '<': operator.lt,
    '>=': operator.ge,
    '<=': operator.le,
    '=': operator.eq,
    '%': operator.mod,
    'begin': begin,
})


def evaluate(expr, env):
    # a
    if isinstance(expr, str):  # variable
        return env[expr]

    # 2.3
    if not isinstance(expr, list):  # constant literal
        return expr

    op, *rest = expr

    # (if (> 1 2) 10 20)
    # TODO: (if (> 1 2) 3)
    if op == 'if':
        test, true_expr, false_expr = rest
        expr = true_expr if evaluate(test, env) else false_expr
        return evaluate(expr, env)

    # (or (> 1 2) (> 3 2)) -> 1.0
    # (or) -> 0.0
    if op == 'or':
        val = 1.0
        for expr in rest:
            val = evaluate(expr, env)
            if val:
                break
        return val

    # (and (> 1 2) (> 3 2)) -> False
    # (and) -> 1.0
    if op == 'and':
        val = 0.0
        for expr in rest:
            val = evaluate(expr, env)
            if not val:
                break
        return val

    # (define x (* y 7)) -> None
    if op == 'define':
        name, child = rest
        val = evaluate(child, env)
        env[name] = val
        return val

    # (set! x (* y 7)) -> None
    if op == 'set!':
        name, child = rest
        for m in env.maps:
            if name in m:
                val = evaluate(child, env)
                m[name] = val
                return val
        raise NameError(name)

    # (lambda (n) (+ n 1)) -> Lambda
    if op == 'lambda':
        args, body = rest
        return Lambda(args, body, env)

    # (* 3 7)
    func = evaluate(op, env)
    args = [evaluate(arg, env) for arg in rest]
    val = func(*args)
    if val in (True, False):  # Convert True/False from gt, lt and friends
        val = 1.0 if val else 0.0
    return val


# REPL = read, eval, print, loop
def repl():
    while True:
        try:
            code = input('» ')
            if not code.strip():
                continue
            expr = reader(code)
            val = evaluate(expr, builtin)
            if val is not None:
                print(lispify(val))
        except (EOFError, KeyboardInterrupt):
            print('Yalla Bye ☺')
            return
        except Exception as err:
            print('\033[31m', end='')  # red
            print(f'[ERROR] {err.__class__.__name__}: {err}')
            print('\033[0m', end='')  # reset color


def lispify(val):
    if not isinstance(val, list):
        return str(val)

    return '(' + ' '.join(lispify(v) for v in val) + ')'


if __name__ == '__main__':
    import readline  # arrows, history, paren matching ... # noqa

    print('Welcome to humble lisp (hit CTRL-C or CTRL-D to exit)')
    repl()
