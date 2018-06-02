from argparse import ArgumentParser, FileType
from time import sleep
from tqdm import tqdm


iter_tasks = range


def process(n):
    sleep(0.01)


parser = ArgumentParser(description=__doc__)
parser.add_argument(
    'input', help='input file', type=FileType('r'),
    default='-', nargs='?')
args = parser.parse_args()

for task in tqdm(iter_tasks(1000)):
    process(task)
