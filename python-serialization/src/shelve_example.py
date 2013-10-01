import shelve
from pickle_example import p1, p2

db = shelve.open('/tmp/looney.shelve')  # <1>
db['p1'] = p1
db['p2'] = p2
db.close()

db = shelve.open('/tmp/looney.shelve')
assert len(db.keys()) == 2, 'missing data'
assert db['p2'].x == p2.x, 'corrupted data'
