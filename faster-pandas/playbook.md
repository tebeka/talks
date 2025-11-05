activate venv
maximize terminal
bigger font

%run setup.ipy

mention using ipython (like jupyer) and several command line

!glow means.md

We'll talk about CPU & Memory

!glow -w0 rules.md

!figlet -k 'Why Not?'

- 350M instructions per blink (times 12 cores = 4.2B)
- hard to maintain
    - bugs
- precision
- development effort
- risk

!figlet -k 'Why Yes?'
- $$$
    - cloud cost
    - happy users (100ms)
    - No need to move to new tech
- Fun, need to know a lot

Talk on requirements

# A request should return in less than 100ms 99.9% of the time

!glow fool.md

story on getting right data for benchmarks

!pygmentize imports.py
%run imports.py


!cowsay CPU

NYC Taxi data
%time df = pd.read_csv('yellow_tripdata_2020-03.csv.gz')
%time df = pd.read_csv('yellow_tripdata_2020-03.csv.gz', engine='pyarrow')
%time df = pd.read_csv('yellow_tripdata_2020-03.csv.gz', engine='pyarrow', dtype_backend='pyarrow')
%time df = pd.read_parquet('yellow_tripdata_2020-03.parquet')
%time df = pd.read_parquet('yellow_tripdata_2020-03.parquet', dtype_backend='pyarrow')

!viu burn.png
no spec, comma in data, everything is string, can't jump to middle

%timeit?

import lzma
with lzma.open('log.txt.xz', 'rt') as fp:
    lines = fp.readlines()

len(lines)
lines[7]

!pygmentize parse_log.py
%run parse_log.py


%%timeit
df = pd.DataFrame()
for line in lines:
    log = parse_line(line)
    ldf = pd.DataFrame(log, index=[0])
    df = pd.concat([df, ldf], ignore_index=True)


Say it ran once, can tweak

joke on warm you at winter

# pytest-benchmark

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
!viu snakeviz.png

OR

%%prun -D parse.pprof
df = pd.DataFrame()
for line in lines:
    log = parse_line(line)
    ldf = pd.DataFrame(log, index=[0])
    df = pd.concat([df, ldf], ignore_index=True)
!snakeviz parse.pprof


%%timeit
dfs = []
for line in lines:
    log = parse_line(line)
    ldf = pd.DataFrame(log, index=[0])
    dfs.append(ldf)
df = pd.concat(dfs, ignore_index=True)


%timeit df = pd.DataFrame.from_records(parse_line(line) for line in lines)

calculate time diff

df = pd.read_parquet('yellow_tripdata_2020-03.parquet')
%timeit max(df['total_amount'])
%timeit df['total_amount'].max()
%timeif df['total_amount'].values.max()


calculate time diff


s = pd.Series([1, np.nan, 3])
s.max()
s.values.max()

No?
    df['VendorID'].sample(5)
    !glow -w0 vendor.md
    df['VendorID'].unique()

names = {
    1: 'Creative',
    2: 'Curb',
    6: 'Myle',
    5: 'Helix',
}

%timeit df['vendor'] = df['VendorID'].apply(lambda vid: names.get(vid))
%timeit df['vendor'] = df['VendorID'].apply(names.get)
%timeit df['vendor'] = df['VendorID'].map(names)

if you're 1% better every day, year
    1.01**365

size = 50_000
bids = pd.DataFrame({
    'client_1': np.random.randint(1, 1000, size),
    'client_2': np.random.randint(1, 1000, size),
    'client_3': np.random.randint(1, 1000, size),
})

%%timeit
total = 0
for _, row in bids.iterrows():
    total += row.max()


%timeit bids.apply(np.max, axis=1).sum()
%timeit bids.apply(np.max, axis=1, raw=True).sum()

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

!glow -w0 rules.md
    show 3

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

pd.read_parquet?
    (filters)
    SQL

df['total_amount'].max()
f'{Out[68]:,}

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
!glow guide.md

!figlet -f big -c 'Culture >> Process'

!glow -w0 links.md
!viu promo.png
