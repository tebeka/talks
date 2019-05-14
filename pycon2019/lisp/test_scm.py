import scm
import pytest


cases = [
    ('collatz.scm', 1, '(collatz 7)', 22.0),
    ('fact.scm', 1, '(fact 10)', 3628800.0),
    ('circle.scm', 2, '(circle-area 10)', 314.159265358),
    ('square.scm', 1, '(square 7)', 49),
    ('account.scm', 2, '(acct -32)', 68),
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

    out = scm.run(expr, env)
    assert abs(out - expected) < 0.001, expr
