import json
import sys

for line in sys.stdin:
    obj = json.loads(line)
    print(obj)
