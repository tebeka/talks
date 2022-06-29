import json
from dataclasses import dataclass, asdict
from datetime import datetime, timezone
from base64 import b64encode, b64decode

def default(obj):
    if isinstance(obj, datetime):
        return obj.isoformat()
    if isinstance(obj, bytes):
        return b64encode(obj).decode('ascii')
    return obj


def obj_hook(obj):
    obj['created'] = datetime.fromisoformat(obj['created'])
    obj['icon'] = b64decode(obj['icon'])
    return obj

@dataclass
class User:
    id: str
    login: str
    created: datetime
    icon: bytes = b''

    def to_json(self):
        return json.dumps(asdict(self), default=default)

    def validate(self):
        if not self.id:
            raise ValueError('missing id')
        if not self.login:
            raise ValueError('missing login')
        now = datetime.now(tz=timezone.utc)
        if self.created > now:
            raise ValueError('created is in the future')

    @classmethod
    def from_json(cls, data):
        kw = json.loads(data, object_hook=obj_hook)
        u = cls(**kw)
        u.validate()
        return u


u7 = User(
    id='007',
    login='Bond',
    created=datetime(1953, 3, 13, tzinfo=timezone.utc),  # Casino Royale
    icon=b'\x89PNG\r\n\x1a\x0a...',
)
print('u7:', u7)


data = u7.to_json()
print('data:', data)
print('back:', User.from_json(data))


# Moonraker
data = '{"id":"","login":"M","created":"2955-04-05T00:00:00+00:00"}'
#m = User.from_json(data)
#


q = User(
    id='81',
    login='Q',
    created=datetime(1964, 9, 22, tzinfo=timezone.utc),  # Goldfinger
    icon=b'\x89PNG\r\n\x1a\x0a...',
)

users = [
    u7,
    q,
]

# See also HTTP chunked transfer encoding
from socket import socketpair
w, r = socketpair()


print('m:', m)
