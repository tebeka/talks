# Y3K

import pandas as pd

y3k = pd.Timestamp(3000, 1, 1)
print(f'They arrived on {y3k:%B %d}.')
