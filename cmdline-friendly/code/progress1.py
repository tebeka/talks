from argparse import ArgumentParser, FileType
from time import sleep
from itertools import cycle


def process_line(line):
    sleep(0.1)


parser = ArgumentParser(description=__doc__)
parser.add_argument(
    'input', help='input file', type=FileType('r'),
    default='-', nargs='?')
args = parser.parse_args()

spinner = cycle(r'-\|/')
for line in args.input:
    c = next(spinner)
    print(f' {c}\r', end='')
    process_line(line)
