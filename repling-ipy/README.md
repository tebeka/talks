# REPLing with IPython (actually Jupyter Console)

Given at [PyWeb-IL](https://www.meetup.com/PyWeb-IL/) on May 2017

## History & Misc
* In, Out
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
* %edit
* %env
* %run, %run -n
* %hist -o -n -f hist.log
* %logstart -ro, %logstop
* %notebook hist.ipynb
* %pdb
* %prun
* %pwd
* %reset
* %store a, %store -r
* %who, %whos
* %%time for multi line cell

## Config

* jupyter console --generate-config
* confirm_exit
* editing_mode
* highlighting_style    
