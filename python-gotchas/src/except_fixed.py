try:
    raise ValueError
except (IndexError, ValueError) as e:  # <1>
    print('oops')
# => oops
