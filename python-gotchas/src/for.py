from items import items

for count, _ in enumerate(items, 1):
    pass

print('There are {} items'.format(count))
# Traceback (most recent call last):
#   File "src/for.py", line 6, in <module>
#     print('There are {} items'.format(count))
# NameError: name 'count' is not defined


