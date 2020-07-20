#!/usr/bin/env python

import pandas as pd

s = pd.Series([1, 2, 3])
s += pd.Series([4, 5])
print(s)
