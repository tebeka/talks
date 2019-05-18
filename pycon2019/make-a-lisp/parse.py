import ast


def print_tree(node, depth=0):
    print(' ' * 4 * depth, end='')
    print(node.__class__.__name__)
    for child in ast.iter_child_nodes(node):
        print_tree(child, depth+1)


def print_ast(code):
    print(code)
    node = ast.parse(code)
    print_tree(node)
