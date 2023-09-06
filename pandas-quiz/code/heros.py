# Holding Out for a Hero

import pandas as pd

heros = pd.Series(['Batman', 'Wonder Woman', 'Superman'])
found = heros.str.find('Iron Man').any()
print('Wrong universe' if found else 'DC')
