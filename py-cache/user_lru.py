from dataclasses import dataclass
from time import sleep

from functools import lru_cache

@dataclass
class User:
    login: str
    name: str
    roles: list[str]


@lru_cache(1024)
def get_user(login):
    sleep(0.1)  # simulate DB access
    return User(login, 'Bruce', ['ADMIN'])

