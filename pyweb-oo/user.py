import json


class User:
    def __init__(self, login, groups):
        self.name = login
        self.groups = groups

    @classmethod
    def from_json(cls, data):
        attrs = json.loads(data)
        return cls(**attrs)


class Admin(User):
    pass


json_data = '''
{
    "login": "elliot", 
    "groups": ["docker", "wheel", "elliot"]
}
'''
