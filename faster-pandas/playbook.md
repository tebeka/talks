## Rules
`!glow rules.md`
- first performance goals
- percentile
- talk on same hardware & data

    !lscpu
    !free -g

## Code

```python
# %%

import pandas as pd
import numpy as np

pd.__version__

df = pd.read_parquet('yellow_tripdata_2022-04.parquet')
len(df)
format(len(df), ',')
df.columns

# %%
%cow CPU


%time _ = pd.read_parquet('yellow_tripdata_2022-04.parquet')
%time _ = pd.read_csv('yellow_tripdata_2022-04.csv.gz')
%time _ = pd.read_csv('yellow_tripdata_2022-04.csv.gz', engine='pyarrow')
%time _ = pd.read_csv('yellow_tripdata_2022-04.csv.gz', engine='pyarrow', dtype_backend='pyarrow')

# %%

%time df['total_amount'].sum()
%timeit df['total_amount'].sum()

# %%

%%timeit
totals = {}
for vid in df['VendorID'].unique():
    totals[vid] = df[df['VendorID'] == vid]['total_amount'].sum()

# %%

%%prun
totals = {}
for vid in df['VendorID'].unique():
    totals[vid] = df[df['VendorID'] == vid]['total_amount'].sum()

# line_profiler, snakeviz, py-spy, ...

# %%
%timeit df.groupby('VendorID')['total_amount'].sum()

# %%
df10k = df[:10_000]

%%timeit
total = 0
for _, row in df10k.iterrows():
    if row['VendorID'] == 2:
        total += row['total_amount']

# %%
%timeit total = df10k[df10k['VendorID'] == 2]['total_amount'].sum()

# note units

# %%

%timeit max(df['total_amount'])
%timeit df['total_amount'].max()
%timeit df['total_amount'].values.max()

# %%

s = pd.Series([1, np.nan, 3])
s.sum()
s.values.sum()

# pd has many nan's

%timeit df[(df['VendorID'] == 2) & (df['passenger_count'] > 1) & (df['trip_distance'] > 2)]
%timeit df.query('VendorID == 2 & passenger_count > 1')


# %%
## Memory

%cow Memory

memory slide

mb = 1<<20
df.memory_usage().sum() / mb

# %%
amt_df = pd.read_parquet('yellow_tripdata_2022-04.parquet', columns=['VendorID', 'total_amount'])
amt_df.memory_usage().sum() / md;lijb

# %%
df['total_amount'].memory_usage() / mb
df['total_amount'].describe()
np.finfo(np.float32)
df['total_amount'].astype(np.float32).memory_usage() / mb

# %%
names = {
    1: 'Creative',
    2: 'VeriFone',
}

%timeit df['vendor'] = df['VendorID'].apply(lambda vid: names.get(vid))
%timeit df['vendor'] = df['VendorID'].apply(names.get)
%timeit df['vendor'] = df['VendorID'].map(names)


# %%
df['VendorID'].memory_usage()
df['vendor'].memory_usage()
df['vendor'].memory_usage(deep=True)

# %%

s = df['vendor'].astype('string[pyarrow]')
s.memory_usage(deep=True)/df['vendor'].memory_usage(deep=True)

# %%
df['vendor'] = df['vendor'].astype('category')
df['vendor'].memory_usage(deep=True)

# Slides - Talk on process
# pytest-benchmark


---
Backup

numba & Cython

python -m cProfile bench_tokenize.py
python -m cProfile -s cumulative bench_tokenize.py
python -m cProfile -o tokenize.pprof bench_tokenize.py

# %%
import nlp
text = 'In the face of ambiguity, refuse the temptation to guess.'

%timeit nlp.tokenize(text)

nlp_opt.py

%timeit nlp_opt.tokenize(text)
