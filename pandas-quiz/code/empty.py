# Full of It

import pandas as pd

empty = pd.Series([], dtype='float64')
print('full' if empty.all() else 'empty')
