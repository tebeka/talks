# A Delicious Div Sum

import pandas as pd

v1 = pd.Series([0, 2, 4])
v2 = pd.Series([0, 1, 2])
out = v1 // v2
print(out.sum())
