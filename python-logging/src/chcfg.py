#!/usr/bin/env python

import struct, socket
from logsrv import port

with open('src/log.conf') as fo:
    data = fo.read().replace('level = ERROR', 'level = INFO')  # <1>

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(('localhost', port))
sock.send(struct.pack('>L', len(data)) + data)  # <2>
sock.close()
