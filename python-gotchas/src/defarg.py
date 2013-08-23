def append_print(i, items=[]):
    items.append(i)
    print(items)

append_print(1)
# => [1]
append_print(2)
# => [1, 2]
