"""HTTP client using sockets"""
from socket import socket


request_template = '''\
GET {path} HTTP/1.1
Host: {host}
Connection: close

'''

MB = 1 << 20


def get(host, path, max_size=1*MB):
    request = request_template.format(host=host, path=path)

    sock = socket()  # By default TCP socket
    with sock:
        sock.connect(('www.353solutions.com', 80))
        sock.sendall(request.encode())

        buf = []  # Using [].append and ''.join is much faster than data += ...
        size = 0
        while True:
            data = sock.recv(1024)
            if not data:
                break
            buf.append(data)
            size += len(data)
            # Don't read more than max_size into memory
            if size > max_size:
                raise ValueError('reply too large')

    data = b''.join(buf)
    return data.decode('utf-8')
