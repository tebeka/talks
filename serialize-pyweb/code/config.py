import sys

if sys.platform == 'win32':
    max_sockets = 100
else:
    max_sockets = 1000


port = 9021
auth_servers = [
    'auth1.local',
    'auth2.local',
]


del sys  # Cleanup
