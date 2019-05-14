import re
import operator
from collections import ChainMap
from dataclasses import dataclass


@dataclass
class Lambda:
    args: list
    body: list
    env: dict

    def __call__(self, *params):
        args = {name: val for name, val in zip(self.args, params)}
        env = ChainMap(args, self.env)
        return evaluate(self.body, env)

    def __repr__(self):
        args = ' '.join(self.args)
        body = lispify(self.body)
        return f'(lambda ({args}) {body})'


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
        tokens.pop(0)  # remove closing )
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


def parse(code):
    # TODO: Multiple s-expressions
    tokens = tokenize(code)
    return read_sexpr(tokens)


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
})


def evaluate(expr, env):
    if isinstance(expr, str):  # variable
        return env[expr]

    if not isinstance(expr, list):  # constant literal
        return expr

    op, *rest = expr

    if op == 'if':
        test, true_expr, false_expr = rest
        expr = true_expr if evaluate(test, env) else false_expr
        return evaluate(expr, env)

    if op == 'or':
        for expr in rest:
            if evaluate(expr, env):
                return True
        return False

    if op == 'and':
        for expr in rest:
            if not evaluate(expr, env):
                return False
        return True

    if op == 'begin':
        x = None
        for child in rest:
            x = evaluate(child, env)
        return x

    if op == 'define':
        name, child = rest
        val = evaluate(child, env)
        env[name] = val
        return val

    if op == 'set!':
        name, child = rest
        for m in env.maps:
            if name in m:
                m[name] = evaluate(child, env)
                return
        raise NameError(name)

    if op == 'lambda':
        args, body = rest
        return Lambda(args, body, env)

    func = evaluate(expr[0], env)
    args = [evaluate(arg, env) for arg in expr[1:]]
    return func(*args)


def run(code, env=None):
    env = builtin if env is None else env
    return evaluate(parse(code), env)


def repl():
    while True:
        try:
            code = input('> ')
            val = run(code)
            if val is not None:
                print(lispify(val))
        except (EOFError, KeyboardInterrupt):
            print('ciao')
            return
        except Exception as err:
            print(f'error: {err}')


def lispify(val):
    if not isinstance(val, list):
        return str(val)

    return '(' + ' '.join(lispify(v) for v in val) + ')'
