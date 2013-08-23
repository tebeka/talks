try:
    int('')
except ValueError as e:
    print('error: <context> {}'.format(e))
    raise

# error: <context> invalid literal for int() with base 10: ''
# Traceback (most recent call last):
#   File "src/reraise.py", line 2, in <module> <1>
#     int('')
# ValueError: invalid literal for int() with base 10: ''
