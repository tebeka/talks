class CustomError(Exception):
    def __init__(self, origin):
        super(CustomError, self).__init__(*origin.args)
        self.origin = origin

try:
    int('')
except ValueError as e:
    print('error: {}'.format(e)) # <1> Use original exception
    raise CustomError(e)
# error: invalid literal for int() with base 10: ''
# Traceback (most recent call last):
#   File "src/other2.py", line 10, in <module> <2>
#     raise CustomError(e)
# __main__.CustomError: invalid literal for int() with base 10: ''
