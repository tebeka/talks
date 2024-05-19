import json

loc = (32.071517, 34.8445123)
data = json.dumps(loc)
loc2 = json.loads(data)
print(loc == loc2)  # False
print(loc)
# (32.071517, 34.8445123)
print(loc2)
# [32.071517, 34.8445123]
