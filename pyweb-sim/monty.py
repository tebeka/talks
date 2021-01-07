"""Simulation of "Monty Hall problem"

See https://en.wikipedia.org/wiki/Monty_Hall_problem for a description of the
problem.
"""
from random import randint


def random_door():
    """Return random door."""
    return randint(1, 3)


def stay_win_game():
    """Simulate a game and returns True if the say policy wins."""
    car_door = random_door()
    player_door = random_door()
    return car_door == player_door  # True if stay wins


n = 1_000_000
stay_wins, switch_wins = 0, 0
for _ in range(n):
    if stay_win_game():
        stay_wins += 1
    else:
        switch_wins += 1

stay_frac = stay_wins / n
switch_frac = switch_wins / n
print(f'stay: {stay_frac:.2f}')
print(f'switch: {switch_frac:.2f}')
