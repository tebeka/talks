import attr


@attr.s
class Node:
    value = attr.ib()
    left = attr.ib()
    right = attr.ib()

    def insert(self, value):
        attr = 'left' if value < self.value else 'right'
        obj = getattr(self, attr)
        node = obj.insert(value) if obj else Node(value, None, None)
        setattr(self, attr, node)
        return self


def sum_tree(tree):
    if tree is None:
        return 0
    return tree.value + sum_tree(tree.left) + sum_tree(tree.right)


if __name__ == '__main__':
    from random import shuffle, seed

    seed(42)
    tree = Node(0, None, None)
    values = list(range(1, 10))
    shuffle(values)
    for val in values:
        tree.insert(val)

    print(sum_tree(tree))
