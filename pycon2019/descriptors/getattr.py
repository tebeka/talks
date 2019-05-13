def get_attr(obj, name):
    if name in obj.__dict__:
        print('found in object')
        return obj.__dict__[name]
    for cls in obj.__class__.__mro__:
        if name in cls.__dict__:
            print(f'found in class {cls.__name__}')
            return cls.__dict__[name]
    raise AttributeError(f'{type(obj)} object has no attribute {name!r}')
