try:
    int('10')
except ValueError as e:
    print('error: {}'.format(e))
else:
    print('OK')  # <1>
