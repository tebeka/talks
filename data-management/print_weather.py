#!/usr/bin/env python

import sqlite3
import pandas as pd

conn = sqlite3.connect('weather.db')
df = pd.read_sql('SELECT * FROM weather', conn, parse_dates=['DATE'])
df = df[['DATE', 'SNOW', 'TMAX', 'TMIN', 'PGTM']]
print(df.head())
