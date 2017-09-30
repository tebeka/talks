from joblib import Memory
import pandas as pd

from os.path import expanduser

memory = Memory(cachedir=expanduser('~/.cache/myapp'), verbose=0)


@memory.cache
def load(filename):
    df = pd.read_csv(filename, parse_dates=['start_time', 'end_time'])
    # ...
    return df


@memory.cache
def pre_process(df):
    df['temp'].fillna(0, inplace=True)
    # ...



