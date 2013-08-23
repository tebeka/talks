from items import items

count = 0  # <1>
for count, _ in enumerate(items, 1):
    pass
print('There are {} items'.format(count))

# OR

count = sum(1 for i in items)  # <2>
print('There are {} items'.format(count))
