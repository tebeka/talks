# Y10K

import pandas as pd

y10k = pd.Timestamp(10_000, 1, 1)
print(f'They arrived on {y10k:%A}.')
