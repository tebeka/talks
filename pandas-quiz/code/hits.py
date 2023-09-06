# Hit and Run

import sqlite3
import pandas as pd

conn = sqlite3.connect(':memory:')
conn.executescript('''
CREATE TABLE visits (day DATE, hits INTEGER);
INSERT INTO visits VALUES
    ('2020-07-01', 300),
    ('2020-07-02', 500),
    ('2020-07-03', 900);
''')

df = pd.read_sql('SELECT * FROM visits', conn)
print('time span:', df['day'].max() - df['day'].min())
