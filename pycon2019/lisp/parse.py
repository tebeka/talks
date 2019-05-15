import ast
from textwrap import dedent


def print_ast(node, depth=0):
    print(' ' * 4 * depth, end='')
    print(node.__class__.__name__)
    for child in ast.iter_child_nodes(node):
        print_ast(child, depth+1)


def demo():
    code = '''
    if x > 10:
        x /= 2
    '''

    m = ast.parse(dedent(code))
    print(code)
    print_ast(m)
