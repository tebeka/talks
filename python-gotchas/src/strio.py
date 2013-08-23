from io import StringIO  # <1>

stream = StringIO('Thank you Sir, may I have another?')
# Traceback (most recent call last):
#   File "src/strio.py", line 3, in <module>
#     stream = StringIO('Thank you Sir, may I have another?')
# TypeError: initial_value must be unicode or None, not str
