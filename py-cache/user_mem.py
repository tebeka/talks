from dataclasses import dataclass
from time import sleep

from joblib import Memory

memory = Memory('/tmp/users', verbose=False)

@dataclass
class User:
    login: str
    name: str
    roles: list[str]


@memory.cache
def get_user(login):
    sleep(0.1)  # simulate DB access
    return User(login, 'Bruce', ['ADMIN'])

