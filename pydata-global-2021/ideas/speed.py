from time import perf_counter
import sqlite3


schema_no_check = '''
CREATE TABLE numbers (
    n INTEGER
)
'''

schema_check = '''
CREATE TABLE numbers (
    n INTEGER CHECK(typeof(n) == 'integer')
)
'''

insert_sql = 'INSERT INTO numbers VALUES (?)'


def timeit(schema, nums):
    conn = sqlite3.connect(':memory:')
    conn.executescript(schema)
    conn.commit()
    start = perf_counter()
    with conn:
        conn.executemany(insert_sql, nums)
    t = perf_counter() - start
    cur = conn.execute('SELECT MAX(n) FROM numbers')
    n = cur.fetchone()[0]
    assert n == nums[-1][0]
    return t


nums = [(n,) for n in range(1_000_000)]
t = timeit(schema_no_check, nums)
print(f'no check: {t}')

t = timeit(schema_check, nums)
print(f'check   : {t}')
