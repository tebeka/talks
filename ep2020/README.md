# IPython: The Productivity Booster

Given at [EuroPython 2020](https://ep2020.europython.eu/talks/5LGWwvT-ipython-the-productivity-booster/), July 2017

## Talk

### Setup

- terminator
- zsh -f
- PS1='$ '
- clear & banner
- start logger
<!-- - `make run-docker` -->

### Why?
- Work in REPL - productivity
    - Testing at first stages too restrictive

### History & Misc
```
me = 'Miki Tebeka <miki@353solutions.com>'
927 - 574
n = Out[1]
f'miki@{n}solutions.com'

from pathlib import Path
Path.<TAB>

passwds = {'daffy': 'Rabbit Season', 'bugs': 'Duck Season'}
passwds['d<TAB>

Path.is_dir?
?Path.isdir
Path.is_dir??
```

### Shell
```
log_dir = '/var/log'
!ls $log_dir
!ls {log_dir}
class config:
    log_dir = '/var/log'
!ls {config.log_dir}
def log_dir():
    return '/var/log'
!ls {log_dir()}

files = !ls -lth /var/log
files, len(files), files[-1]
type(files)
files.grep('Xorg')
files.grep(lambda v: 'xo' in v.lower())
files.fields(5, 6, 7)

!xzcat track.csv.bz2 | head | xsel -b
df = pd.read_clipboard(sep=',', parse_dates=['time'])
```

### Magic
- %pwd
    - %pwd?
- %time
- %timeit
    %timeit?
    %timeit '%s' % 'hi'
    %timeit '{}'.format('hi')
- %edit
    %edit t.py
    %edit and then %edit -p
    %edit -x
- %run, %run -n
    %run t.py
    %run -n t.py
- %env
    %env HOME
    %env DEBUG true
- %who, %whos
- %reset
    x = 10
    %reset
    x
- %store a, %store -r
    a = 10
    %store a
    exit
    %store -r
    a
<!--
- %hist -o -n -f hist.log
- %logstart -ro, %logstop
- %notebook hist.ipynb   < FIXME >
    jupyter notebook
-->
- %pdb
    %edit -x t.py
    def div(a, b):
	return a / b
    div(1/2)
    %run t.py
    %pdb
    %run t.py
- %prun
    %edit nlp.py
    import this (pick one)
    s = '...'
    %prun tokenize(s)
    %%prun
    for _ in range(100_000):  # joke on time left
	tokenize(s)
<!--
    %edit httpc.py
    %prun get('www.353solutions.com', '/')
- %%time for multi line cell

- matplotlib qt5 < FIXME >
    - Do without before
-->

- load track
- df - see lat/lng
- pd.options.display.max_rows = 5
- pd.options.display.float_format = '{:.3f}'.format
- df - see lat/lng
- mention df.style in Jupyter Notebook

- %load_ext sql
    - pip install ipython-sql
    - https://github.com/catherinedevlin/ipython-sql
- %sql sqlite:///weather.db
- %sql SELECT MIN(TMAX), MAX(TMAX) FROM weather
    - same with %%sql
- result = %sql SELECT * FROM weather WHERE TMAX > 0
- df = result.DataFrame()
    - %config SqlMagic
    - %config SqlMagic.autopandas = True
- async def inc(n): return n + 1
- await inc(10)
    - !python
    - show the above fails (python -m asyncio?)
- %autoawait

### Config

- %edit -x ~/.ipython/profile_default/ipython_config.py
- jupyter console --generate-config
- ipython profile create
- confirm_exit
- editing_mode
- highlighting_style    
- %config
- autoreload
```
c.InteractiveShellApp.exec_lines = [
    '%autoreload 2',
]
### A list of dotted module names of IPython extensions to load.
c.InteractiveShellApp.extensions = [
    'autoreload',
]
```
- from hello import hello
- hello()
- %edit -x hello.py
- hello()

### IDE Integration
- PyCharm
    - Use IPython
    - Run selection in console
- VSCode
    - 

---
Comments
- More humor / war stories
- Story?
