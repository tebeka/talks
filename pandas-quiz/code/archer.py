# A Late Addition

import pandas as pd

df = pd.DataFrame([
    ['Sterling', 83.4],
    ['Cheryl', 97.2],
    ['Lana', 13.2],
], columns=['name', 'sum'])
df.late_fee = 3.5
print(df)
