import numpy as np
import pandas as pd

# Initial look at the data

csv_file = 'taxi.csv.bz2'
# !ls -lh $csv_file
# !bzcat $csv_file | wc -l
# !bzcat $csv_file | head


# Load taxi.csv.bz2, make sure time columns are parsed as time


time_cols = ['tpep_pickup_datetime', 'tpep_dropoff_datetime']
df = pd.read_csv(csv_file, parse_dates=time_cols)
df.dtypes


# How many trips per VendorID?
df['VendorID'].value_counts()
df['VendorID'].unique()

# apply example
s = pd.Series([1, 2, 3, 4])



# Create a column called 'Vendor', populate from VendorID
# - 1 -> Creative
# - 2 -> VeriFone
# - 4 -> BigApple (until we find out)
# Use apply


# %%timeit

names = {
    1: 'Creative',
    2: 'VeriFone',
    4: 'BigApple',
}
df['Vendor'] = df['VendorID'].apply(names.get)
df['Vendor'].sample(5)


# %%timeit
df['Vendor'] = df['VendorID'].map(names)
df['Vendor'].sample(5)


# UPDATE df ADD COLUMN Vendor FROM
# SELECT CASE
#   WHEN VendorID == 1 THEN 'Creatvie'
#   WHEN VendorID == 2 THEN 'VeriFone'
#   When VendorID == 4 THEN 'BigApple'
# END AS Vendor


# Categorical data
df.memory_usage().sum() / 2**20  # MB
df['Vendor'] = df['Vendor'].astype('category')
df.memory_usage().sum() / 2**20  # MB
df['Vendor'].cat.categories
df['Vendor'].sample(6) == 'VeriFone'

# Create a box plot (.plot.box()) of trip time in minutes
# make the y axis lograithmic

t1 = pd.Timestamp('2019-04-07T15:53')
t2 = pd.Timestamp('2019-04-07T15:43')
dt = t1 - t2
# dt.total_seconds() / 60
dt / np.timedelta64(1, 'm')

duration = df['tpep_dropoff_datetime'] - df['tpep_pickup_datetime']
duration /= np.timedelta64(1, 'm')  # Convert to minutes
duration.head()
duration.describe()
# %matplotlib  notebook
clean = duration[(duration > 0) & (duration < 1440)]  # Remove some outliers
ax = clean.plot.box(logy=True, label='Trip Time')
ax.set_ylabel('Duration (min)')  # ;


# How many rides between 11am and 12pm had more than 1 passanger?
len(df[
    (df['tpep_pickup_datetime'].dt.hour == 11) &
    (df['passenger_count'] > 1)])

len(df.query('tpep_pickup_datetime.dt.hour==11 & passenger_count > 1'))


# How many rides of VeriFone + BigApple
len(df[df['Vendor'].isin({'VeriFone', 'BigApple'})])

# --- day 3 ---
df.groupby('Vendor')  # DataFrameGroupBy
for k, sdf in df.groupby('Vendor'):
    print(f'{k} -> {len(sdf)}')

c = df.groupby('Vendor').count()
c
c.index
df.groupby('Vendor')['VendorID'].count()
# SELECT Vendor, COUNT(VendorID) FROM df GROUP BY Vendor

# Which day of week has longer rides?
df.groupby(df['tpep_pickup_datetime'].dt.weekday)['trip_distance'].median()
# SELECT WEEKDAY(tpep_pickup_datetime) day, MEDIAN(trip_distance) distance
# FROM df
# GROUP BY WEEKDAY(tpep_pickup_datetime)

# Create a bar chart of median number of rides per hour. Split the chart by
# VendorID
df['hour'] = df['tpep_pickup_datetime'].dt.hour
df['day'] = df['tpep_pickup_datetime'].dt.date
counts = df.groupby(['day', 'hour'], as_index=False).count()
# SELECT MEDIAN(count) FROM counts GROUP BY hour
hourly = counts.groupby('hour').median()['Vendor']  # Any column will do
hourly.plot.bar(rot=45)
# TODO: customize

# Earnings per day per vendor
sub = df[['Vendor', 'total_amount', 'day']]
psub = sub.pivot_table(
    index='day', columns='Vendor', values='total_amount', aggfunc=np.sum)

# Tip according to weather
# https://rp5.ru/Weather_archive_in_New_York,_La_Guardia_(airport)
# Do people tip more on cold days?
wdf = pd.read_csv(
    'data/nyc_weather.csv.gz', skiprows=6, sep=';',
    parse_dates=['Local time in New York / Downtown Manhattan'])
wdf['day'] = wdf['Local time in New York / Downtown Manhattan'].dt.date
temp = wdf.groupby('day', as_index=False)['T'].mean()
tdf = pd.merge(df, temp, on='day')
tdf['bin'] = pd.cut(tdf['T'], 10)
tdf.groupby('bin')['tip %'].median()

# Categorical data
df.memory_usage().sum() / 2**20  # MB
df['Vendor'] = df['Vendor'].astype('category')
df.memory_usage().sum() / 2**20  # MB
df['Vendor'].cat.categories
df['Vendor'].sample(6) == 'VeriFone'

# Join
temp = wdf.groupby('day', as_index=False)['T'].mean()
tdf = pd.merge(df, temp, on='day')
tdf['bin'] = pd.cut(tdf['T'], 10)
tdf.groupby('bin')['tip %'].median()
