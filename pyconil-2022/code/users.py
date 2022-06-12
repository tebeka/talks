import json
from dataclasses import dataclass, asdict
from datetime import datetime, timezone
from base64 import b64encode, b64decode


@dataclass
class User:
    id: str
    login: str
    created: datetime
    icon: bytes

    def as_dict(self):
        return asdict(self)


u7 = User(
    id='007',
    login='Bond',
    created=datetime(1953, 3, 13, 14, 33, 42, tzinfo=timezone.utc),
    icon=b'\x89PNG\r\n...',
)


def default(obj):
    if isinstance(obj, datetime):
        return obj.isoformat()
    if isinstance(obj, bytes):
        return b64encode(obj).decode('utf-8')
    return obj

data = json.dumps(u7.as_dict(), default=default)
print(data)

request = json.loads(data)
u = User(**request)
print(u)


def obj_hook(obj):
    obj['created'] = datetime.fromisoformat(obj['created'])
    obj['icon'] = b64decode(obj['icon'].encode('utf-8'))
    return obj

request = json.loads(data, object_hook=obj_hook)
u = User(**request)
print(u)

uQ = User(
    id='81',
    login='Q',
    created=datetime(1964, 9, 22, 17, 32, 44, tzinfo=timezone.utc),
    icon=b'\x89PNG\r\n...',
)

users = [
    u7,
    uQ,
]

from socket import socketpair

rs, ws = socketpair()
r = rs.makefile('rb')
w = ws.makefile('wb')

for u in users:
    data = json.dumps(u.as_dict(), default=default)
    data = data.encode('utf-8')
    print(b'1', file=w)
