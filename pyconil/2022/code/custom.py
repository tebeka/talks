from datetime import datetime, timezone
import json


def default(obj):
    if isinstance(obj, datetime):
        return obj.isoformat()
    return obj


event = {
    'time': datetime.now(tz=timezone.utc),
    'type': 'click',
    'user': 'joe',
}
data = json.dumps(event, default=default)

event2 = json.loads(data)
print(event2['time'])
# 2022-06-12T13:39:02.639146+00:00


def obj_hook(obj):
    obj['time'] = datetime.fromisoformat(obj['time'])
    return obj


event2 = json.loads(data, object_hook=obj_hook)
print(event['time'] == event2['time'])  # True
