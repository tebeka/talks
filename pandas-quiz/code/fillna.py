# Phil? Nah!?

import numpy as np
import pandas as pd

values = pd.Series([1, np.nan, 3])
values.fillna(2)
print(values.sum())
