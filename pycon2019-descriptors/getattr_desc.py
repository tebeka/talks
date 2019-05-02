def get_attr(obj, name):
    if name in obj.__dict__:
        return obj.__dict__[name]
    for cls in obj.__class__.__mro__:
        if name in cls.__dict__:
            val = cls.__dict__[name]
            if hasattr(val, '__get__'):
                return val.__get__(obj, cls)
            return val
    raise AttributeError(f'{type(obj)} object has no attribute {name!r}')
