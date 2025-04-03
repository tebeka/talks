# Identify the User

next_uid = 1


class User:
    def __init__(self, name):
        global next_uid

        self.name = name
        self.__id = next_uid
        next_uid += 1


u = User('daffy')
print(f'name={u.name}, id={u.__id}')
