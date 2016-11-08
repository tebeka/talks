#!/usr/bin/env python

import pandas as pd
import sqlite3
from os import remove


def read_csv(path):
    return pd.read_csv(path, dtype={'DATE': str})


df = pd.concat([
    read_csv('564288.csv'),
    read_csv('564289.csv'),
    pd.read_csv('564290.csv'),
])

dbfile = 'weather.db'

df = df[df['STATION'] == 'GHCND:USW00094846']
try:
    remove(dbfile)
except OSError:
    pass
conn = sqlite3.connect(dbfile)
df.to_sql('weather', conn)
