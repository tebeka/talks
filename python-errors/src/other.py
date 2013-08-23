class CustomError(Exception):
    pass

try:
    int('')
except ValueError as e:
    print('error: {}'.format(e))
    raise CustomError('some message')

# error: invalid literal for int() with base 10: ''
# Traceback (most recent call last):
#   File "src/other.py", line 8, in <module> <1>
#     raise CustomError('some message')
# __main__.CustomError: some message
