import aiohttp
from http import HTTPStatus

api_url = 'http://localhost:9876/check'


async def check_ip(ip):
    """Check that IP is not malicious"""
    params = {
        'ip': ip,
    }

    async with aiohttp.ClientSession() as session:
        async with session.get(api_url, params=params) as resp:
            if resp.status != HTTPStatus.OK:
                raise ValueError('bad ip')
            resp = await resp.json()

    return resp['ok']
