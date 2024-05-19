import scm
import pytest


cases = [
    ('collatz.scm', 1, '(collatz 7)', 22.0),
    ('fact.scm', 1, '(fact 10)', 3628800.0),
    ('circle.scm', 2, '(circle-area 10)', 314.159265358),
    ('square.scm', 1, '(square 7)', 49),
    ('account.scm', 2, '(acct -32)', 68),
    ('adder.scm', 2, '(add-7 10)', 17),
    ('', 0, '(- 1 7)', -6),
]

env0 = scm.builtin.copy()


@pytest.mark.parametrize('fname, nsexpr, expr, expected', cases)
def test_scm(fname, nsexpr, expr, expected):
    env = env0.copy()
    if fname:
        with open(fname) as fp:
            code = fp.read()

        tokens = scm.tokenize(code)
        for i in range(nsexpr):
            ast = scm.read_sexpr(tokens)
            scm.evaluate(ast, env)

    out = scm.evaluate(scm.reader(expr), env)
    assert abs(out - expected) < 0.001, expr


logic_cases = [
    ('(or)', 1.0),
    ('(or 1 2)', 1.0),
    ('(or 0 2 1)', 2.0),
    ('(and)', 0.0),
    ('(and 1 2)', 2.0),
    ('(and 1 0 3)', 0.0),
]


@pytest.mark.parametrize('expr, expected', logic_cases)
def test_logic(expr, expected):
    env = env0.copy()
    out = scm.evaluate(scm.reader(expr), env)
    assert out == expected, 'bad logic'
