# Rectified!

import pandas as pd


def relu(n):
    if n < 0:
        return 0
    return n


arr = pd.Series([-1, 0, 1])
print(relu(arr))
