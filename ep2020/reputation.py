#!/usr/bin/env python

# IP data from https://gist.github.com/bsmartt13/efa02c40ea12c09d9c3a

import gzip
from http import HTTPStatus
from ipaddress import ip_address
from pathlib import Path

from fastapi import FastAPI, HTTPException

db = None
app = FastAPI()


def load_db(filename):
    ips = set()
    with gzip.open(filename, 'rt') as fp:
        for line in fp:
            line = line.strip()
            if not line or line[:1] == '#':
                continue
            i = line.find(' ')
            if i == -1:
                continue
            ips.add(line[:i])
    return ips


@app.get('/check')
async def check_ip(ip):
    try:
        ip_address(ip)
    except ValueError:
        raise HTTPException(HTTPStatus.BAD_REQUEST, f'invalid ip: {ip!r}')

    return {
        'ip': ip,
        'ok': ip not in db,
    }


@app.on_event('startup')
def on_startup():
    global db

    here = Path(__file__).absolute().parent
    db_file = here / 'reputation.generic.gz'
    db = load_db(db_file)


if __name__ == '__main__':
    import uvicorn

    uvicorn.run(app, port=9876)
