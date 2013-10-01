import csv
tmp = '/tmp/points.csv'
points = [{'x': x, 'y': x+1} for x in range(10)]

with open(tmp, 'w') as out:
    writer = csv.DictWriter(out, ['x', 'y'])
    writer.writeheader()
    writer.writerows(points)

with open(tmp) as fo:
    reader = csv.DictReader(fo)
    points1 = [point for point in reader]  # Values are strings

def apply_schema(obj, schema):
    return dict((key, schema.get(key, lambda x: x)(value))
                for key, value in obj.iteritems())

points2 = [apply_schema(point, {'x': int, 'y': int}) for point in points1]
assert points2 == points, 'corrupted data'
