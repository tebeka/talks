"""Spam with n Messages"""
from argparse import ArgumentParser

parser = ArgumentParser(description=__doc__)
parser.add_argument('count', type=int, help='number of messages')
args = parser.parse_args()

for i in range(args.count):
    print(f'Spam message #{i+1}')
