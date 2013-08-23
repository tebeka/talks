#!/usr/bin/env python3

import nums

print(nums.a)
# => 1
print(nums.b)
# => 2

print(nums.b < nums.a)
# Traceback (most recent call last):
#   File "src/compare_fixed.py", line 8, in <module>
#     print(nums.b < nums.a)  # => True
# TypeError: unorderable types: int() < str()
