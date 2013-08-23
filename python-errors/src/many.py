try:
    int('')
except IndexError as e:
    print('error: IndexError - {}'.format(e))
except ValueError as e:
    print('error: ValueError - {}'.format(e))
