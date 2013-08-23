try:
    int('')
except ValueError as e:
    print('error: {}'.format(e))
finally:
    print('finally')  # <1>
