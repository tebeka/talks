#!/usr/bin/env python

from argparse import ArgumentParser
from sys import stdin

parser = ArgumentParser()
parser.add_argument('first')
parser.add_argument('last')
args = parser.parse_args()

n = 1
should_change = False
prefix = '## '
plen = len(prefix)
for line in stdin:
    if line.startswith(prefix) and args.first in line:
        should_change = True

    if args.last in line:
        should_change = False

    if not should_change or not line.startswith(prefix):
        print(line, end='')
        continue

    if should_change or args.first in line:
        should_change = True
    else:
        continue

    print(f'## [{n}] {line[plen:]}', end='')
    n += 1
