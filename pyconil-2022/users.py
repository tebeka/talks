import json
from dataclasses import dataclass, asdict
from datetime import datetime, timezone


def default(obj):
    if isinstance(obj, datetime):
        return obj.isoformat()
    return obj


@dataclass
class User:
    id: str
    login: str
    created: datetime

    def to_json(self):
        return json.dumps(asdict(self), default=default)

    @classmethod
    def from_json(cls, data):
        d = json.loads(data)
        u = cls(**d)
        return u

u7 = User(
    id='007',
    login='Bond',
    created=datetime(1953, 3, 13, tzinfo=timezone.utc),  # Casino Royale
)
print('u7:', u7)


data = u7.to_json()
print('data:', data)
u = User.from_json(data)
print(u)
