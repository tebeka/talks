import json


class User:
    def __init__(self, login, groups):
        self.name = login
        self.groups = groups


class Admin(User):
    pass


json_data = '''
{
    "login": "elliot", 
    "groups": ["docker", "wheel", "elliot"]
}
'''
