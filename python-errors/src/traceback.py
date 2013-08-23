def inner(v):
    return int(v)
def outer(v):
    return inner(v)
outer('')
# Traceback (most recent call last): <1>
#   File "src/stacktrace.py", line 7, in <module>
#     outer('')
#   File "src/stacktrace.py", line 5, in outer <2>
#     return inner(v)
#   File "src/stacktrace.py", line 2, in inner
#     return int(v)
# ValueError: invalid literal for int() with base 10: '' <3>
