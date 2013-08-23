def append_print(i, items=None):
    items = items or []
    items.append(i)
    print(items)

append_print(1)
# => [1]
append_print(2)
# => [2]
