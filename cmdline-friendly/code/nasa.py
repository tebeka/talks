"""Which hour of day is the busiest according to NASA logs?

Keep just one line in memory at all times.

Bonus: If you can't parse a line - print the file and line where the error is

Hints: glob.iglob, datetime.strptime, collections.Counter
"""

# In [307]: ts = '01/Aug/1995:00:00:08'
#
# In [308]: datetime.strptime(ts, '%d/%b/%Y:%H:%M:%S')
# Out[308]: datetime.datetime(1995, 8, 1, 0, 0, 8)

from collections import Counter
from datetime import datetime
from glob import iglob
from itertools import repeat, count
from operator import attrgetter
from sys import stderr


def iter_lines(pattern):
    for path in iglob(pattern):
        with open(path) as fp:
            for lnum, line in enumerate(fp, 1):
                yield from zip(repeat(path), count(1), fp)


def parse_date(record):
    name, lnum, line = record
    # uplherc.upl.com - - [01/Aug/1995:00:00:07 -0400] "GET ...
    fields = line.split()
    try:
        time_stamp = fields[3][1:]  # Remove [ prefix
        return datetime.strptime(time_stamp, '%d/%b/%Y:%H:%M:%S')
    except (ValueError, IndexError):
        print(f'{name}:{lnum}: warning - cannot parse {line!r}', file=stderr)
        return None


lines = iter_lines(f'nasa-logs/*.log')
times = filter(None, map(parse_date, lines))
hours = map(attrgetter('hour'), times)
counts = Counter(hours)
for hour, value in counts.most_common(3):
    print(f'{hour}: {value}')
