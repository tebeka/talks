import json
from base64 import b64decode, b64encode
from dataclasses import asdict, dataclass
from datetime import datetime, timedelta, timezone
from socket import socketpair


@dataclass
class User:
    id: str
    login: str
    created: datetime
    icon: bytes = b''

    def to_json(self):
        return json.dumps(asdict(self), default=default)

    @classmethod
    def from_json(cls, data):
        kw = json.loads(data, object_hook=obj_hook)
        u = cls(**kw)
        u.validate()
        return u

    def validate(self):
        if not self.id:
            raise ValueError('missing id')
        if not self.login:
            raise ValueError('missing login')
        now = datetime.now(tz=timezone.utc)
        if self.created > now + timedelta(hours=1):
            raise ValueError('created is in the future')


u7 = User(
    id='007',
    login='Bond',
    created=datetime(1953, 3, 13, tzinfo=timezone.utc),
    icon=b'\x89PNG\r\n\x1a\x0a...',
)
print('u7:', u7)


def default(obj):
    if isinstance(obj, datetime):
        return obj.isoformat()
    if isinstance(obj, bytes):
        return b64encode(obj).decode('utf-8')
    return obj


data = u7.to_json()
print('data:', data)

def obj_hook(obj):
    obj['created'] = datetime.fromisoformat(obj['created'])
    if 'icon' in obj:
        obj['icon'] = b64decode(obj['icon'].encode('utf-8'))
    return obj

u = User.from_json(data)
print('u:', u)




request = json.loads(data, object_hook=obj_hook)
u = User(**request)
print('u:', u)


data = '{"id":"","login":"M","created":"2955-04-05T00:00:00+00:00"}'
request = json.loads(data, object_hook=obj_hook)
m = User(**request)
# m.validate()
print('m:', m)

q = User(
    id='81',
    login='Q',
    created=datetime(1964, 9, 22, tzinfo=timezone.utc),
    icon=b'\x89PNG\r\n\x1a\x0a...',
)

users = [
    u7,
    q,
]


w, r = socketpair()

for u in users:
    data = u.to_json()
    data = data.encode('utf-8')
    w.sendall(data + b'\n')
w.close()

# print(r.recv(1024).decode('utf-8'))

fp = r.makefile('r')
for line in fp:
    u = User.from_json(line)
    print('got:', u)
