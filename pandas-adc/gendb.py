#!/usr/bin/env python

from zipfile import ZipFile
import sqlite3
from os.path import isfile
from os import remove
import pandas as pd
import numpy as np

def sql_vals(values):
    return tuple(int(val) if type(val) == np.int64 else val for val in values)


def insert_from(db, table, fo):
    df = pd.read_csv(fo, sep='\1')
    names = ', '.join(df.columns)
    qmarks = ', '.join(['?'] * len(df.columns))

    sql = 'INSERT INTO {} ({}) VALUES ({})'.format(
        table, names, qmarks)

    for row in df.itertuples(index=False):
        values = sql_vals(row)
        db.execute(sql, values)


def main():
    zf = ZipFile('data.zip')
    dbfile = 'edw.sqlite'
    if isfile(dbfile):
        remove(dbfile)

    db = sqlite3.connect(dbfile)

    with open('schema.sql') as fo:
        db.executescript(fo.read())

    # Populate facts
    for name in zf.namelist():
        if '2013' in name:
            insert_from(db, 'FACT_OMNITURE_PAGE_VIEWS', zf.open(name))

    # Populate DIM_SITE
    insert_from(db, 'DIM_SITE', zf.open('dim_site.csv'))

    db.commit()


if __name__ == '__main__':
    main()
