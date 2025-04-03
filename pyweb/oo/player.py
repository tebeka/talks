# Ready Player One


class Player:
    # Number of players in the Game
    count = 0

    def __init__(self, name):
        self.name = name
        Player.count += 1


p1 = Player('Parzival')
print(Player.count)
