# IPython: The Productivity Booster

Given at [EuroPython 2020](https://ep2020.europython.eu/talks/5LGWwvT-ipython-the-productivity-booster/), July 2017

## History & Misc
- In, Out
- me = 'miki@353solutions.com'
- Completion `os.<TAB>`, `d['x<TAB>`
- path.isdir?, path.isdir??, ?path.isdir


## Shell
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

## Magic
- %pwd
- %time
- %timeit
    %timeit?
    %timeit '%s' % 'hi'
    %timeit '{}'.format('hi')
- %edit
    %edit t.py
    %edit and then %edit -p
- %env
    %env HOME
    %env TMPDIR /tmp
- %run, %run -n
    %run t.py
    %run -n t.py
- %hist -o -n -f hist.log
- %logstart -ro, %logstop
- %notebook hist.ipynb
    jupyter notebook
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
<!--
    %edit httpc.py
    %prun get('www.353solutions.com', '/')
-->
- %pwd
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
- %who, %whos
- %%time for multi line cell
- matplotlib qt5
    - Do without before

- load track
- df - see lat/lng
- pd.options.display.max_rows = 5
- pd.options.display.float_format = '{:.2f}'.format
- df - see lat/lng

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

## Config

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
## A list of dotted module names of IPython extensions to load.
c.InteractiveShellApp.extensions = [
    'autoreload',
]
```
- from hello import hello
- hello()
- %edit -x hello.py
- hello()
