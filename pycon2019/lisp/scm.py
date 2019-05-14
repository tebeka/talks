import re
import operator
from collections import ChainMap
from dataclasses import dataclass


@dataclass
class Function:
    args: list
    body: list
    env: dict

    def __call__(self, *params):
        args = {name: val for name, val in zip(self.args, params)}
        env = ChainMap(args, self.env)
        return evaluate(self.body, env)


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


builtin = {
    '+': operator.add,
    '-': operator.sub,
    '*': operator.mul,
    '/': operator.truediv,
    '>': operator.gt,
    '<': operator.lt,
    '>=': operator.ge,
    '<=': operator.le,
    '=': operator.eq,
    'even?': lambda x: x % 2 == 0,
}


def evaluate(expr, env):
    if isinstance(expr, str):  # variable
        return env[expr]

    if not isinstance(expr, list):  # constant literal
        return expr

    if expr[0] == 'if':
        _, test, true_expr, false_expr = expr
        expr = true_expr if evaluate(test, env) else false_expr
        return evaluate(expr, env)

    if expr[0] == 'or':
        for expr in expr[1:]:
            if evaluate(expr, env):
                return True
        return False

    if expr[0] == 'and':
        for expr in expr[1:]:
            if not evaluate(expr, env):
                return False
        return True

    if expr[0] == 'begin':
        x = None
        for child in expr[1:]:
            x = evaluate(child, env)
        return x

    if expr[0] == 'define':
        _, name, child = expr
        val = evaluate(child, env)
        env[name] = val
        return val

    if expr[0] == 'lambda':
        _, args, body = expr
        return Function(args, body, env)

    func = evaluate(expr[0], env)
    args = [evaluate(arg, env) for arg in expr[1:]]
    return func(*args)


def ep(code):
    return evaluate(parse(code), builtin)
