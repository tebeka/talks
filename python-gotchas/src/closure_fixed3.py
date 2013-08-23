#!/usr/bin/env python3

def make_counter():
    x = 0
    def counter():
        nonlocal x  # <1>
        x += 1
        return x
    return counter

c = make_counter()
print(c())
# => 1
print(c())
# => 2
