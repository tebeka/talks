def make_counter():
    x = 0
    def counter():
        x += 1
        return x
    return counter

c = make_counter()
print(c())
# Traceback (most recent call last):
#   File "src/closure.py", line 10, in <module>
#     print(c())
#   File "src/closure.py", line 4, in counter
#     x += 1
# UnboundLocalError: local variable 'x' referenced before assignment

