# We Should Have Turned Left At Albuquerque

MAX_X, MAX_Y = 100, 100


def advance(point: tuple):
    """Advance point one step up to (MAX_X, MAX_Y)"""
    match point:
        case (MAX_X, MAX_Y):
            return (MAX_X, MAX_Y)
        case (x, MAX_Y):
            return (x + 1, MAX_Y)
        case (MAX_X, y):
            return (MAX_X, y+1)
        case (x, y):
            return (x+1, y+1)

point = (5, 100)
print(advance(point))
