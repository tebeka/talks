import sqlite3

points = [{'x': x, 'y': x+1} for x in range(10)]  # <1>

db = sqlite3.connect(':memory:')
db.row_factory = sqlite3.Row  # <2>
with db:
    cur = db.cursor()
    cur.execute('CREATE TABLE points (x INT, y INT)')
    cur.executemany('INSERT INTO points (x, y) VALUES (:x, :y)', points)  # <3>

cur.execute('SELECT x, y FROM points ORDER BY x, y')
points1 = [dict(row) for row in cur]

assert points1 == points, 'corrupted data'
