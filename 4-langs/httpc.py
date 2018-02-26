from socket import socket
from functools import partial

host = 'www.353solutions.com'

payload = f'''\
GET / HTTP/1.1
Host: {host}
Connection: close

'''

with socket() as sock:
    sock.connect((host, 80))
    sock.sendall(payload.encode('utf-8'))

    data = b''.join(iter(partial(sock.recv, 1024), b''))


print(data.decode('utf-8')[:400])
