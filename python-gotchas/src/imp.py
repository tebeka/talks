import email
with open('src/clp.eml') as input:
    message = email.message_from_file(input)
# Traceback (most recent call last):
#   File "src/imp.py", line 2, in <module>
#     message = email.message_from_file('/tmp/1.eml')
# AttributeError: 'module' object has no attribute 'message_from_file' <1>

