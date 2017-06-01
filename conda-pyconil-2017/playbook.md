* http://mutt.org/ - credits
* What's wrong with pip? Let's SSH to a machine and try to install `pyodbc`

    $ docker run --rm -it pycoil

You didn't think I'll anger the demo gods with going to the internet :)

    $ which python

"The first rule of Python is you don't use system python." - Kelsey Hightower

I actually have a Python from python.org install at /opt as well. So you want
say I'm cheating :)

    $ /opt/bin/pip3 install pyodbc

We need `sql.h`, we need a C compiler ...
Even when we manage to build, we can install on another machine and find out
we're missing system libraries and then install them with `apt`, `yum`,
`pacman` ...

To be fair, `manylinux1` distribution is a big help here. But it uses CentOS 5
as the build system. You're missing about 10 years of compiler advances, bug
fixes ...

    $ conda install pyodbc

Note the we got unixodbc as well.

    $ ls /miniconda/lib/*odbc*
    $ exit

OK. So how do we work with conda?

We have two options, install Anaconda or Miniconda. Anaconda comes bundled with
many packages and unpacked is about 1.9GB. Miniconda comes with Python and
conda and is about 134MB unpacked.

On linux you download a .sh file and then execute. It even has a batch mode.


    $ docker run -it --rm -v ${HOME}/Downloads/:/dl pyconil-ub
    $ bash /dl/Miniconda3-latest-Linux-x86_64.sh -p /miniconda -b
    $ export PATH=/miniconda/bin:$PATH
