activate venv
maximize terminal

!clear
!glow welcome.md
!glow means.md

!glow rules.md

Why not
- 350M instructions per blink (times 12 cores = 4.2B)
- hard to maintain
- bugs

Why yes
- $$$ (backend, users 100ms)
- No need to move to new tech

!glow fool.md

story on getting right data for benchmarks

!pygmentize imports.py
%run imports.py

!cowsay CPU

%time df = pd.read_csv('yellow_tripdata_2020-03.csv.gz')
%time df = pd.read_csv('yellow_tripdata_2020-03.csv.gz', engine='pyarrow')
%time df = pd.read_csv('yellow_tripdata_2020-03.csv.gz', engine='pyarrow', dtype_backend='pyarrow')
%time df = pd.read_parquet('yellow_tripdata_2020-03.parquet')
%time df = pd.read_parquet('yellow_tripdata_2020-03.parquet', dtype_backend='pyarrow')

!viu burn.png

%timeit?

!pygmentize parse_log.py
%run parse_log.py

import lzma
with lzma.open('log.txt.xz', 'rt') as fp:
    lines = fp.readlines()

%%timeit
df = pd.DataFrame()
for line in lines:
    log = parse_line(line)
    ldf = pd.DataFrame(log, index=[0])
    df = pd.concat([df, ldf], ignore_index=True)


joke on warm you at winter

pytest-benchmark

!pygmentize test_logs.py
!python -m pytest -v


%prun?

%%prun -s cumulative
df = pd.DataFrame()
for line in lines:
    log = parse_line(line)
    ldf = pd.DataFrame(log, index=[0])
    df = pd.concat([df, ldf], ignore_index=True)

!glow profs.md


%%timeit
dfs = []
for line in lines:
    log = parse_line(line)
    ldf = pd.DataFrame(log, index=[0])
    dfs.append(ldf)
df = pd.concat(dfs, ignore_index=True)


%%timeit
df = pd.DataFrame.from_records(parse_line(line) for line in lines)


calculate time diff


%timeit max(df['total_amount'])
%timeit df['total_amount'].max()

arr = df['total_amount'].to_numpy()
%timeit arr.max()

calculate time diff


s = pd.Series([1, np.nan, 3])
s.max()
s.values.max()

!glow -w0 vendor.md
df['VendorID'].unique()

names = {
    1: 'Creative',
    2: 'Curb',
    6: 'Myle',
    5: 'Helix',  # 7 in schema
}

%timeit df['vendor'] = df['VendorID'].apply(lambda vid: names.get(vid))
%timeit df['vendor'] = df['VendorID'].apply(names.get)
%timeit df['vendor'] = df['VendorID'].map(names)


size = 50_000
bids = pd.DataFrame({
    'client_1': np.random.randint(1, 1000, size),
    'client_2': np.random.randint(1, 1000, size),
    'client_3': np.random.randint(1, 1000, size),
})

%%timeit
total = 0
for _, row in df.iterrows():
    total += row.max()


%timeit df.apply(np.max, axis=1).sum()
%timeit df.apply(np.max, axis=1, raw=True).sum()

def relu(v):
    if v < 0:
        return 0
    return v

s = pd.Series(np.random.randint(-3, 200, 1_000_000))
%timeit s.apply(relu)

import numba
@numba.vectorize
def vect_relu(n):
    if n <= 0:
        return 0
    return n

%timeit vect_relu(s)

calculate diffs

!cowsay Memory

why is memory important
- cost
- latency

!viu virtual-memory.png
!glow latency.md

mb = 1<<20
df.memory_usage(deep=True).sum() / mb

from pathlib import Path
Path('yellow_tripdata_2020-03.parquet').stat().st_size / mb

amt_df = pd.read_parquet(
    'yellow_tripdata_2020-03.parquet', 
    columns=['VendorID', 'total_amount'],
    dtype_backend='pyarrow',
)
amt_df.memory_usage(deep=True).sum() / mb

pd.read_parquet? (filters)
SQL

df['total_amount'].max()

import pyarrow.parquet as pq
max_amount = 0
pfile = pq.ParquetFile('yellow_tripdata_2020-03.parquet')
for batch in pfile.iter_batches():
    bdf = batch.to_pandas()
    if (m := bdf['total_amount'].max()) > max_amount:
        max_amount = m
max_amount

can't batch everything (median)


df['vendor'].memory_usage(deep=True) / mb
df.dtypes

s = df['vendor'].astype('string[pyarrow]')
s.memory_usage(deep=True) / mb

c = df['vendor'].astype('category')
c.memory_usage(deep=True) / mb

!glow -w0 flow.md

!figlet -f big -c 'Culture >> Process'

!viu promo.png
