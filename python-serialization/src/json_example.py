import json

data = {
    'name': 'Daffy',
    'type': 'Duck',
    'age': 76,
}

tmp = '/tmp/looney.json'
with open(tmp, 'wb') as out:
    json.dump(data, out)

with open(tmp) as fo:
    data1 = json.load(fo)

assert data1 == data, 'corrupted data'
