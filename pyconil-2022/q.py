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

from socket import socketpair
w, r = socketpair()
