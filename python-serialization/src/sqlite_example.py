import sqlite3

points = [(x, x+1) for x in range(10)]

db = sqlite3.connect(':memory:')  # <1>
cur = db.cursor()
with db:  # <2>
    cur.execute('CREATE TABLE points (x INT, y INT)')
    cur.executemany('INSERT INTO points (x, y) VALUES (?, ?)', points)

cur.execute('SELECT x, y FROM points ORDER BY x, y')
points1 = list(cur)  # <3>

assert points1 == points, 'corrupted data'
