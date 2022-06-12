import json

import numpy as np

pi = np.float64(3.14)
print(json.dumps(pi))  # 3.14
n = np.int64(353)
print(json.dumps(n))
