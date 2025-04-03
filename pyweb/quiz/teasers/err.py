# To Err is Human

err = None

try:
    1/0
except ZeroDivisionError as err:
    pass

print(f'ERROR: {err}')
