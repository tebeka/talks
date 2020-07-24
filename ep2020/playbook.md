# Setup

- python reputation.py (different shell)
- terminator
- start logger
- ipython
- Open PyCharm on settings, have logs.py open
- `In [] me = 'Miki Tebeka <miki@353solutions.com>'

# Why?

- Work in REPL - productivity
- Testing at first stages too restrictive
- PG code sketches
- No context switch to another app

- # Task: Load compressed log files from a directory to a DataFrame

Magic commands, start with %, not valid Python

    >>> %pwd
    '/home/miki/Projects/talks/ep2020'
    >>> %magic

Output history, if you forget to capture in out

    >>> logs_dir = Out[1] + '/logs'

! for calling commands, can pass variables using $

    >>> !ls $logs_dir
    >>> !ls {logs_dir}

Can assign output to a variables

    >>> files = !ls -l {logs_dir}
    >>> types(files)
    >>> files.<TAB>  # see attributes
    >>> files.g<TAB> # attributes starting with g
    >>> log_files = files.grep('log')

Size of all files

    >>> sum(int(val) for val in log_files.fields(4))

In MB

    >>> Out[]/(1<<20)
Or we can use du

    >>> !du -sh $logs_dir

Pick a log file

    >>> log_file = logs_dir + '/' + log_files.fields(-1)[0]

Use xzcat to get lines, later we'll use lzma

    >>> lines = !xzcat $log_file | head -50
    >>> lines
    >>> line = lines[0]
    >>> line
    >>> line.split()
    >>> fields = Out[]
    >>> for i, v in enumerate(fields): 
    ...:     print(f'{i}: {v}') 

Talk on fields
Copy output

    >>> %edit logs.py  # logs1.py, there's also %run if you edit outside
    logs1.py
    >>> parse_line(line)
    >>> for line in lines:
    ...     print(parse_line(line))
    >>> from pprint import pprint
    >>> for <UP ARROW> # change to pprint
    >>> for line in lines:
    ...     pprint(parse_line(line))
    >>> import lzma

IPython built in help

    >>> lzma.open?  # note the 'rb' - we'll need 'rt'
    >>> %timeit?
    >>> lzma.open??

Forgot the name of the variable for log file
    %who

    >>> with lzma.open(log_file, 'rt') as fp:
    ...     for line in fp:
    ...         parse_line(line)
    >>> %pdb  # text based, highly effectiv3
    >>> with lzma.open(log_file, 'rt') as fp:
    ...     for line in fp:
    ...         parse_line(line)

Talk on pbb
    ipdb> ?


size can be '-'

    >>> %edit logs.py  # logs2.py
    >>> with lzma.open(log_file, 'rt') as fp:
    ...     for line in fp:
    ...         parse_line(line)

    >>> %pdb  # turn off pdb


cell magic

    >>> %timeit parse_line(line)
    >>> %%time
    ... with lzma.open(log_file, 'rt') as fp:
    ...     for line in fp:
    ...         parse_line(line)

Load to data frame

    >>> import pandas as pd
    >>> records = [parse_line(line) for line in lzma.open(log_file, 'rt')]
    >>> df = pd.DataFrame.from_records(records)
    >>> df
    >>> df.iloc[353]
    >>> df.sample(5)
    >>> df

TMI

    >>> pd.options.display.max_rows = 5
    >>> df

Many option, float formatting. There's style for jupyter notebooks
We'll stop here - you can finish the code
    >>> pd.describe_options('display')

    >>> %history -p -o -f hist.log

New requirement - merge with weather data in sqlite database

    >>> # %pip install ipython-sql
    >>> %load_ext sql
    >>> %sql sqlite:///weather.db
    >>> %sql select name, sql from sqlite_master
    >>> print(Out[])
    >>> %sql SELECT MIN(TMAX) AS min_temp, MAX(TMAX) as max_temp FROM weather
    >>> # same with %%sql
    >>> result = %sql SELECT * FROM weather WHERE TMAX > 0
    >>> result
    >>> wdf = result.DataFrame()
    >>> %config SqlMagic
    >>> %config SqlMagic.autopandas = True
    >>> result = %sql SELECT * FROM weather WHERE TMAX > 0
    >>> result
    >>> # continue with pandas read_sql

New task: Ad a column to check that origin is valid (TIME)

    >>> %edit check_ip.py
    >>> check_ip('8.8.8.8')
    couroutine
    >>> await check_ip('8.8.8.8')

It's magic! IPython works as expected

    >>> !python
    >>> from check_ip import check_ip
    >>> await check_ip('8.8.8.8')
    >>> # SyntaxError, use python -m asyncio
    >>> CTRL-D # exit Python

- PyCharm
    - Use IPython
    - Run selection in console

You can also configuration IPython. For example nag on exit, highlight style
autoreload

    >>> %config
    >>> %edit -x ~/.ipython/profile_default/ipython_config.py

Thanks

    >>> print(me)
    >>> # https://github.com/tebeka/talks/tree/master/ep2020
    >>> # Python Brain Teasers (https://gum.co/iIQT)
