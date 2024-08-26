from dataclasses import dataclass
from time import sleep

from cache import cached

@dataclass
class User:
    login: str
    name: str
    roles: list[str]


@cached
def get_user(login):
    sleep(0.1)  # simulate DB access
    return User(login, 'Bruce', ['ADMIN'])

