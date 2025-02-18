"""Simulation of the "Birthday problem"

See https://en.wikipedia.org/wiki/Birthday_problem for a description of the
problem.

Data from https://github.com/fivethirtyeight/data/tree/master/births
"""

import pandas as pd
import numpy as np


def load_birthdays(csv_file):
    """ "Load probability of birthdays from csv_file. Return days,
    probabilities."""
    df = pd.read_csv(csv_file)
    df.rename(columns={'date_of_month': 'day'}, inplace=True)
    times = pd.to_datetime(df[['year', 'month', 'day']])
    df['bday'] = times.dt.strftime('%m-%d')
    probs = df.groupby('bday')['births'].sum()
    probs /= probs.sum()
    return probs.index, probs.values


def dup_in_group(size, days, probs):
    """Return True if there's at least one duplicate birthday in a random
    group of size.
    """
    group = np.random.choice(days, size, p=probs)
    return np.unique(group).size < group.size


days, probs = load_birthdays('US_births_2000-2014_SSA.csv')
n, group_size, dups = 100_000, 23, 0
for _ in range(n):
    if dup_in_group(group_size, days, probs):
        dups += 1

frac = dups / n
print(f'{frac:.2f}')
