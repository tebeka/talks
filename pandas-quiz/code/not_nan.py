# Off With Their NaNs

import numpy as np
import pandas as pd

values = pd.Series([1, np.nan, 3])
print(values[~(values == np.nan)])
