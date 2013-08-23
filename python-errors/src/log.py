import logging as log

try:
    int('')
except ValueError as e:
    log.exception('hosed inting an empty string')  # <1>

# ERROR:root:hosed inting an empty string <2>
# Traceback (most recent call last): <3>
#   File "src/log.py", line 4, in <module>
#     int('')
# ValueError: invalid literal for int() with base 10: ''
