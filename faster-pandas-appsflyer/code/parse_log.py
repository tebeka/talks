"""Parse log lines"""

from datetime import datetime


def parse_time(ts):
    # [02/Jul/1995:16:30:08 -0400]
    time = datetime.strptime(ts, '[%d/%b/%Y:%H:%M:%S %z]')
    return time.replace(tzinfo=None)  # Remove time zone


def parse_line(line):
    fields = line.split()
    size = 0 if fields[-1] == '-' else int(fields[-1])
    return {
        'origin': fields[0],
        'time': parse_time(fields[3] + ' ' + fields[4]),
        'method': fields[5][1:],  # Remove leading "
        'path': fields[6],
        'status_code': int(fields[-2]),
        'size': size,
    }
