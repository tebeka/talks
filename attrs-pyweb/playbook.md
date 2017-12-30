Background story: Manage fleet of autonomous cars

---
    class Coordinate:
	def __init__(self, lat, lng):
	    self.lat = lat
	    self.lng = lng

---
    $ ipython
    In [1]: %run gis.py

    In [2]: gcampus = Coordinate(32.0700849,34.7919559)

    In [3]: gcampus
    Out[3]: <__main__.Coordinate at 0x7fb70e988a58>

---
    def __repr__(self):
        return f'{self.__class__.__name__}({self.lat!r}, {self.lng!r})'

---
    In [4]: %run gis.py

    In [5]: gcampus = Coordinate(32.0700849,34.7919559)

    In [6]: gcampus
    Out[6]: Coordinate(32.0700849, 34.7919559)

---

    def distance(coord1, coord2):
	"""Return distance between two coordinates"""
	# TODO: Haversine
	dlat = coord1.lat - coord2.lat
	dlng = coord1.lng - coord2.lng
	return math.sqrt(dlat**2 + dlng**2) * 111

---
    In [17]: %run gis.py 
    In [18]: home = Coordinate(32.5195835,35.013115)
    
---
    In [24]: distance(gcampus, home)
    Out[24]: 55.60649544584651

http://boulter.com/gps/distance/?from=32.5195835%2C35.013115&to=32.0700849%2C34.7919559&units=k


---
Optimize - distance 0 if equal

    def distance(coord1, coord2):
	"""Return distance between two coordinates"""
	if coord1 == coord2:
	    return 0

	# TODO: Haversine
	dlat = coord1.lat - coord2.lat
	dlng = coord1.lng - coord2.lng
	return math.sqrt(dlat**2 + dlng**2) * 111

You don't see any improvement

---

    In [32]: c1 = Coordinate(1, 2)

    In [33]: c1 == c1
    Out[33]: True

    In [34]: c2 = Coordinate(1, 2)

    In [35]: c1 == c2
    Out[35]: False

Need to write `__eq__` - too much code for this. The solution? ...

namedtuple :)

---
    from collections import namedtuple

    Coordinate = namedtuple('Coordinate', 'lat lng')

    In [38]: c1 = Coordinate(1, 2)

    In [39]: c2 = Coordinate(1, 2)

    In [40]: c1
    Out[40]: Coordinate(lat=1, lng=2)

    In [41]: c2
    Out[41]: Coordinate(lat=1, lng=2)

    In [42]: c1 == c2
    Out[42]: True

---
We'll get telemetry from car and need to update locaton

---
And check it out (restart ipython)


    In [3]: c  = Coordinate(32.0700849,34.7919559)

    In [7]: c.lat = 32.08

    ---------------------------------------------------------------------------
    AttributeError                            Traceback (most recent call last)
    <ipython-input-12-de012420864a> in <module>()
    ----> 1 c.lat = 32.08

    AttributeError: can't set attribute

---
Let's use attrs

    import attr

    @attr.s
    class Coordinate:
	lat = attr.ib()
	lng = attr.ib()

- attr.s -> attrs
- attr.ib -> attrib

---

    In [14]: %run gis.py

    In [15]: c = Coordinate(32.0700849,34.7919559)

    In [16]: c
    Out[16]: Coordinate(lat=32.0700849, lng=34.7919559)

    In [17]: c.lat = 32.08

    In [18]: c
    Out[18]: Coordinate(lat=32.08, lng=34.7919559)


    In [28]: c = Coordinate('hi', 'there')

    In [29]: c
    Out[29]: Coordinate(lat='hi', lng='there')

---
    @attr.s
    class Coordinate:
	lat = attr.ib(validator=attr.validators.instance_of(float))
	lng = attr.ib(validator=attr.validators.instance_of(float))

---
    In [30]: %run gis.py
    In [31]: c = Coordinate('hi', 'there')

---
Add height

    @attr.s
    class Coordinate:
	lat = attr.ib(validator=attr.validators.instance_of(float))
	lng = attr.ib(validator=attr.validators.instance_of(float))
	height = attr.ib()

--- 

    In [34]: %run gis.py

    In [35]: c = Coordinate(32.0700849,34.7919559)
    ---------------------------------------------------------------------------
    TypeError                                 Traceback (most recent call last)
    <ipython-input-35-b0265ce7f2a7> in <module>()
    ----> 1 c = Coordinate(32.0700849,34.7919559)

    TypeError: __init__() missing 1 required positional argument: 'height'


---
    @attr.s
    class Coordinate:
	lat = attr.ib(validator=attr.validators.instance_of(float))
	lng = attr.ib(validator=attr.validators.instance_of(float))
	height = attr.ib(default=0)

default can be a factory (for lists ...)
--- 

    In [34]: %run gis.py
    In [35]: c = Coordinate(32.0700849,34.7919559)
    In [38]: c
    Out[38]: Coordinate(lat=32.0700849, lng=34.7919559, height=0)

---
Need to send over as JSON


    In [41]: attr.asdict(c)
    Out[41]: {'height': 0, 'lat': 32.0700849, 'lng': 34.7919559}

    In [42]: d = attr.asdict(c)

    In [43]: Coordinate(**d)
    Out[43]: Coordinate(lat=32.0700849, lng=34.7919559, height=0)

- Also have `astuple`
- Can filter out what is serialized

---

Much more features
- `__attrs_post_init__` 
- `make_class` for dynamic creation (used in creating classes from XML schema)
- frozen
- converters
- metadata on attributes
- auto attributes with types
- ...

See also https://www.python.org/dev/peps/pep-0557/, coming in 3.7 (June 2018)
