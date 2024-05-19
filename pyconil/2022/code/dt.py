import json
from datetime import datetime, timezone

now = datetime.now(tz=timezone.utc)
json.dumps(now)
