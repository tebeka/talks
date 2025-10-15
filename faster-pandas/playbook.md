!clear
!glow welcome.md
!glow means.md

!glow rules.md

why not, why yes

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

%%timeit
totals = {}
for vid in df['VendorID'].unique():
    totals[vid] = df[df['VendorID'] == vid]['total_amount'].sum()


joke on warm winter


%prun?

%%prun -s cumulative
totals = {}
for vid in df['VendorID'].unique():
    totals[vid] = df[df['VendorID'] == vid]['total_amount'].sum()

!glow profs.md

%timeit df.groupby('VendorID')['total_amount'].sum()

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

!cowsay Memory


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


df['vendor'].memory_usage(deep=True) / mb
df.dtypes

s = df['vendor'].astype('string[pyarrow]')
s.memory_usage(deep=True) / mb

c = df['vendor'].astype('category')
c.memory_usage(deep=True) / mb

!glow -w0 flow.md

!figlet big -c 'Culture >> Process'

!viu promo.png
