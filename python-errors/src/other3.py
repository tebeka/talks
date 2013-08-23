class CustomError(Exception): pass
try:
    int('')
except ValueError as e:
    raise CustomError from e # <1>
# error: invalid literal for int() with base 10: ''
# Traceback (most recent call last):
#   File "src/other3.py", line 5, in <module>
#     int('')
# ValueError: invalid literal for int() with base 10: ''
# 
# The above exception was the direct cause of the following exception:
# 
# Traceback (most recent call last):
#   File "src/other3.py", line 8, in <module>
#     raise CustomError from e
# __main__.CustomError
