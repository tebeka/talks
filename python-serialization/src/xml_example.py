import xml.etree.ElementTree as et

tmp = '/tmp/points.xml'
points = [{'x': str(x), 'y': str(x+1)} for x in range(10)]  # <1>

root = et.Element('points')
tree = et.ElementTree(root)
for point in points:
    root.append(et.Element('point', point))

with open(tmp, 'w') as out:
    tree.write(out)

with open(tmp) as fo:
    tree = et.parse(fo)

epoints = tree.iterfind('.//point')  # <2>
points1 = [dict(ep.items()) for ep in epoints]
assert points1 == points, 'corrpted data'

