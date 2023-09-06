# I've Got Chills...

import pandas as pd

v = pd.Series([.1, 1., 1.1])
out = v * v
expected = pd.Series([.01, 1., 1.21])
if (out == expected).all():
    print('Math rocks!')
else:
    print('Please reinstall universe & reboot.')
