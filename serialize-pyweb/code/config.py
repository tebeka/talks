import sys

max_sockets = 100 if sys.platform == 'win32' else 1000


port = 9021
auth_servers = [
    'auth1.local',
    'auth2.local',
]


del sys  # Cleanup
