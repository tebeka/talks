# data from https://github.com/fivethirtyeight/data/tree/master/births
import pandas as pd
import matplotlib.pyplot as plt

df = pd.read_csv('US_births_2000-2014_SSA.csv')
df.rename(columns={'date_of_month': 'day'}, inplace=True)
df = df[df['year'] == 2013]
df.index = pd.to_datetime(df[['year', 'month', 'day']])
ax = df['births'].plot()
ax.set_ylabel('Number of births')
ax.set_title('US Births 2013')
plt.savefig('birth_rate.png')
