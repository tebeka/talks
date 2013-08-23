try:
    raise ValueError
except IndexError, ValueError:
    print('oops')
# Traceback (most recent call last):
#   File "src/catch.py", line 2, in <module>
#     raise ValueError
# ValueError
