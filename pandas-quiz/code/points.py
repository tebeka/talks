# What’s the Points?

import pandas as pd

df = pd.DataFrame(
    [
        ['Gon', 1001],
        ['Killua', 1000],
    ],
    columns=['name', 'points'],
)
print(df.to_csv())
