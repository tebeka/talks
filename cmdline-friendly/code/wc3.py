"""Count words in lines"""
from argparse import ArgumentParser, FileType
import json

parser = ArgumentParser(description=__doc__)
parser.add_argument(
    'input', help='input file', type=FileType('r'),
    default='-', nargs='?')
parser.add_argument(
    '--output', help='input file', type=FileType('w'),
    default='-')
parser.add_argument(
    '--json', help='JSON formatted output',
    action='store_true', default=False)
args = parser.parse_args()


def report_text(name, lnum, count, out):
    print(f'{name}:{lnum}: {count}', file=out)


def report_json(name, lnum, count, out):
    obj = {
        'file': name,
        'line': lnum,
        'count': count,
    }
    json.dump(obj, out)
    out.write('\n')


if args.json:
    report = report_json
else:
    report = report_text

name = args.input.name
for lnum, line in enumerate(args.input, 1):
    count = len(line.split())
    report(args.input.name, lnum, count, args.output)
