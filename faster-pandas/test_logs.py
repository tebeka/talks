import lzma

import pandas as pd

from parse_log import parse_line


def logs2df(lines):
    df = pd.DataFrame()
    for line in lines:
        log = parse_line(line)
        ldf = pd.DataFrame(log, index=[0])
        df = pd.concat([df, ldf], ignore_index=True)

    return df


def test_bench_load(benchmark):
    with lzma.open('log.txt.xz', 'rt') as fp:
        lines = fp.readlines()

    df = benchmark(logs2df, lines)
    assert len(df) == 1_000
