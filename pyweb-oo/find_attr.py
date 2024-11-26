def find_attr(obj, name):
    # First look in instance attributes
    if name in obj.__dict__:
        print(f'found {name} in instance')
        return obj.__dict__[name]

    # Then look at class
    if name in obj.__class__.__dict__:
        print(f'found {name} in class')
        return obj.__class__.__dict__[name]

    # Then go up inheritance tree
    for cls in obj.__class__.__mro__:
        if name in cls.__dict__:
            print(f'found {name} in class {cls.__name__}')
            return cls.__dict__[name]

    raise AttributeError(name)
