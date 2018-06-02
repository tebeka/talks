"""Print numbers n..."""
from argparse import ArgumentParser
from itertools import count

parser = ArgumentParser(description=__doc__)
parser.add_argument('start', type=int, help='number to start')
args = parser.parse_args()

for n in count(args.start):
    print(n)
