items = [str(i) for i in xrange(5)]  # <1>
print(', '.join(items))
# => 0, 1, 2, 3, 4
print(', '.join(items))
# => 0, 1, 2, 3, 4

from itertools import tee
items = (str(i) for i in xrange(5))
i1, i2 = tee(items)  # <2>
print(', '.join(i1))
# => 0, 1, 2, 3, 4
print(', '.join(i2))
# => 0, 1, 2, 3, 4
