import json
from dataclasses import dataclass
from datetime import datetime, timezone


@dataclass
class User:
    id: str
    login: str
    created: datetime

    def to_json(self):
        return json.dumps(self)

    @classmethod
    def from_json(cls, data):
        kw = json.loads(data)
        u = cls(**kw)
        return u


u7 = User(
    id='007',
    login='Bond',
    created=datetime(1953, 3, 13, tzinfo=timezone.utc),  # Casino Royale
)
print('u7:', u7)


data = u7.to_json()
print('data:', data)
print('back:', User.from_json(data))
