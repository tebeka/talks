# REPLing with IPython (actually Jupyter Console)

Given at [PyWeb-IL](https://www.meetup.com/PyWeb-IL/) on May 2017

## History & Misc
* In, Out
* Completion (os.)
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
    In [21]: !bzcat track.csv.bz2 | head | xsel -b
    In [22]: df = pd.read_clipboard()

## Magic
* %pwd
* %time
* %timeit
* %edit
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
