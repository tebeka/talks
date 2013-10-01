import marshal

data = {
    'name': 'Daffy',
    'type': 'Duck',
    'age': 76,
}

tmp = '/tmp/looney.m'
with open(tmp, 'wb') as out:  # <1>
    marshal.dump(data, out)

with open(tmp) as fo:
    data1 = marshal.load(fo)

assert data1 == data, 'corrupted data'
