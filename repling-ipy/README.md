# REPLing with IPython (actually Jupyter Console)

Given at [PyWeb-IL](https://www.meetup.com/PyWeb-IL/) on May 2017

## History & Misc
* In, Out
* me = 'miki@353solutions.com'
* Completion `os.<TAB>`, `d['x<TAB>`
* path.isdir?, path.isdir??, ?path.isdir


## Shell
    log_dir = '/var/log'
    !ls $log_dir
    !ls {log_dir}
    class config:
	log_dir = '/var/log'
    !ls config.log_dir
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
    _exit_code

    !xzcat track.csv.bz2 | head | xsel -b
    df = pd.read_clipboard(sep=',', parse_dates=['time'])

## Magic
* %pwd
* %time
* %timeit
    %timeit '%s' % 'hi'
    %timeit '{}'.format('hi')
* %edit
    %edit t.py
    %edit and then %edit -p
* %env
    %env HOME
    %env TMPDIR /tmp
* %run, %run -n
    %run t.py
    %run -n t.py
* %hist -o -n -f hist.log
* %logstart -ro, %logstop
* %notebook hist.ipynb
    jupyter notebook
* %pdb
    %edit -x t.py
    def div(a, b):
	return a / b
    div(1/2)
    %run t.py
    %pdb
    %run t.py
* %prun
    %edit httpd.py
    %prun get('www.353solutions.com', '/')
* %pwd
* %reset
    x = 10
    %reset
    x
* %store a, %store -r
    a = 10
    %store a
    exit
    %store -r
    a
* %who, %whos
* %%time for multi line cell
* matplotlib qt5
    * Do without before

* load track
* df - see lat/lng
* pd.options.display.max_rows = 10
* pd.options.display.float_format = '{:.2f}'.format
* df - see lat/lng

* %load_ext sql
* %sql sqlite:///weather.db
* %sql SELECT MIN(TMAX), MAX(TMAX) FROM weather
    * same with %%sql
* result = %sql SELECT * FROM weather WHERE TMAX > 0
* df = result.DataFrame()
    * %config SqlMagic
    * %config SqlMagic.autopandas = True

## Config

* jupyter console --generate-config
* ipython profile create
* confirm_exit
* editing_mode
* highlighting_style    
* %config
