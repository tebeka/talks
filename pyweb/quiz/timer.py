#!/usr/bin/env python
"""Timer with progress bar"""

from tqdm import tqdm
from time import sleep
from argparse import ArgumentParser


def positive(val):
    n = int(val)
    if n <= 0:
        raise ValueError(f'expected positive number, got {val}')
    return n


parser = ArgumentParser(description=__doc__)
parser.add_argument('time', help='time in seconds', type=positive)
args = parser.parse_args()


for _ in tqdm(range(args.time), bar_format='{n}/{total}{bar}'):
    sleep(1)
