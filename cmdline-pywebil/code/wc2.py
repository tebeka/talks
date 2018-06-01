"""Count words in lines"""
from argparse import ArgumentParser, FileType

parser = ArgumentParser(description=__doc__)
parser.add_argument(
    'input', help='input file', type=FileType('r'),
    default='-', nargs='?')
parser.add_argument(
    '--output', help='input file', type=FileType('w'),
    default='-')
args = parser.parse_args()

name = args.input.name
for lnum, line in enumerate(args.input, 1):
    count = len(line.split())
    print(f'{name}:{lnum}: {count}', file=args.output)
