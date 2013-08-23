try:
    int('')
except ValueError as e:  # <1>
    print('error: {}'.format(e))
