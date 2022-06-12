import json
from datetime import datetime, timezone
from io import BytesIO


def default(obj):
    if isinstance(obj, datetime):
        return obj.isoformat()
    return obj


metrics = [
    {
        'time': datetime(2022, 3, 15, 13, 25, 39, tzinfo=timezone.utc),
        'name': 'cpu',
        'value': 37.2,
    },
    {
        'time': datetime(2022, 3, 15, 13, 27, 43, tzinfo=timezone.utc),
        'name': 'memory',
        'value': 1182644,
    },
]


io = BytesIO()

for m in metrics:
    print(json.dumps(m, default=default), file=io)


