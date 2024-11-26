class Node:
    def __init__(self, value, next):
        self.value = value
        self.next = next


class Stack:
    def __init__(self):
        self._head = None

    def push(self, value):
        self._head = Node(value, self._head)

    def pop(self):
        if self._head is None:
            raise ValueError('pop from empty stack')

        value = self._head.value
        self._head = self._head.next
        return value

s = Stack()
for c in 'Python':
    s.push(c)
