def or_(a, b):
    if a:
        return True
    return b


True or 1/0
or_(True, 1/0)
