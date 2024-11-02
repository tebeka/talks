# It's a Date!

import pandas as pd

start = (
    pd.Timestamp
    .fromtimestamp(0, 'UTC')
    .strftime('%Y-%m-%dT%H:%M:%S')
)
times = pd.date_range(start=start, freq='ME', periods=2)
print(times)
