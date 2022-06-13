import json
from dataclasses import dataclass
from datetime import datetime, timezone


@dataclass
class User:
    id: str
    login: str
    created: datetime


u7 = User(
    id='007',
    login='Bond',
    created=datetime(1953, 3, 13, tzinfo=timezone.utc),
)


data = json.dumps(u7)
print('data:', data)
