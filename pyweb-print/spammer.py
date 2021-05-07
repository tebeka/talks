"""Spam with n Messages"""
from argparse import ArgumentParser
from time import sleep

parser = ArgumentParser(description=__doc__)
parser.add_argument('count', type=int, help='number of messages')
parser.add_argument('--flush', action='store_true', default=False)
parser.add_argument('--delay', action='store_true', default=False)
args = parser.parse_args()

for i in range(args.count):
    print(f'Spam message #{i+1}', flush=args.flush)
    if args.delay:
        sleep(1)
