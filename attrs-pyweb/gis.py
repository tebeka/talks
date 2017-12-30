import math


if 0:
    class Coordinate:
        def __init__(self, lat, lng):
            self.lat = lat
            self.lng = lng

        def __repr__(self):
            return f'{self.__class__.__name__}({self.lat!r}, {self.lng!r})'


if 0:
    from collections import namedtuple
    Coordinate = namedtuple('Coordinate', 'lat lng')

import attr

@attr.s
class Coordinate:
    lat = attr.ib(validator=attr.validators.instance_of(float))
    lng = attr.ib(validator=attr.validators.instance_of(float))
    height = attr.ib(default=0)


def distance(coord1, coord2):
    """Return distance between two coordinates"""
    if coord1 == coord2:
        return 0

    # TODO: Haversine
    dlat = coord1.lat - coord2.lat
    dlng = coord1.lng - coord2.lng
    return math.sqrt(dlat**2 + dlng**2) * 111
