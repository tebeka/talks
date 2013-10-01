import yaml, datetime

data = {
    'name': 'Daffy',
    'type': 'Duck',
    'born': datetime.date(1937, 4, 17),  # <1>
}

tmp = '/tmp/looney.yaml'
with open(tmp, 'wb') as out:
    yaml.dump(data, out, default_flow_style=False)  # <2>

with open(tmp) as fo:
    data1 = yaml.load(fo)

assert data1 == data, 'corrupted data'
