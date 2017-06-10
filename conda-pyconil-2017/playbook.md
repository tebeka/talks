http://mutt.org/ - credits

What's wrong with pip? Let's SSH to a machine and try to install `pyodbc`

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


    $ docker run -it --rm -v ${PWD}:/dl pyconil-ub

Just one important tweak
    $ set -o vi

    $ bash /dl/Miniconda3-latest-Linux-x86_64.sh -p /miniconda -b
    $ export PATH=/miniconda/bin:$PATH

Let's update the system
    $ conda update --all -y
    $ python --version
    $ conda info

We use conda both to create environment and install packages. If you don't work
in environments (docker for example) you can install directly to the root
environment.

    $ conda create -n pyconil
    $ ls /miniconda/envs/pyconil

Not even python! conda environments system levels, not just Python. Let's add
Python. First we activate the environment
    $ source activate pyconil
    $ which conda

    $ conda install python

You can specify the Python version you want.

    $ which python
    $ python --version
    $ python -c 'import sys; print(sys.prefix)'

And now we can install packages

    $ conda install -y flask

What's nice about these environments is that you can copy them around and
they'll work.

    $ source deactivate
    $ cp -r /miniconda/envs/pycon/ ~/pycon-env
    $ ~/pycon-env/bin/python
    >>> import flask
    flask

Let's get back to our environment

    $ source activate pycon

We can install specific versions, let's see which are available

    $ conda search boto3
    $ conda install -y boto3=1.4.3

It's always a good idea to keep your dependencies in a file and keep it in
version control so we can recreate environments. In the pip world we use a
requirements file. In conda we use an environment file, which is YAML based.

Once you have a working environment, you can export what you have.

    $ conda env export > environment.yml
    $ vim environment.yml

Now let's trip it down. I prefer to leave the top level dependencies. If you
have people working on different OS this will save you some headache.

    dependencies:
    - boto3=1.4.3
    - flask=0.12.2
    - python=3.6.1

And now let's re-create the environment

    $ source deactivate
    $ conda env create -n pycon2 -f environment.yml
    $ source activate pycon2
    $ python -c 'import boto3; print(boto3.__version__)'

We can also clone an existing environment if you want to test a new package.

    $ conda create --name pycon6 --clone pycon

Not all packages hosted on PyPI are available in the default anaconda channel.
If you don't find a package you need on the default channel you have two
options. The first is to find it in a different channel. Let's install fastavro
which deals with the Avro file format.

    $ conda search fastavro

To search in other channels we're going to use the anaconda client. And since it's a global tool we'll install it to the root environment.

    $ conda install -y -n root anaconda-client
    $ which anaconda
    $ anaconda search fastavro | more

A good once to go with is conda-forge which is a community effort to add many
more packages. Let's install from there

    $ conda install --channel conda-forge fastavro

We can add this to environment.yml, either by adding a "channels" section or
prefixing the package with the channel.

    $ vim environment.yml
    channels:
    - defaults
    - conda-forge
    dependencies:
    - boto3=1.4.3
    - fastavro=0.13.0
    - flask=0.12.2
    - python=3.6.1

    $ conda env create -n pycon3 -f environment.yml

By the way - I'm looking for a new maintainer to `fastavro`. If you use avro
and would like to help in open source - let me know. I'll throw in some
mentoring as well.

Sometime you can't even find the package on any channel. conda plays nice with
pip so you can also install packages from PyPI.

    $ anaconda search timeat
    $ conda install timeat
    $ which pip
    $ pip install timeat
    $ which timeat
    $ timeat london

Let's add this to the environment file
    
    $ vim environment.yml
    channels:
    - defaults
    - conda-forge
    dependencies:
    - boto3=1.4.3
    - fastavro=0.13.0
    - flask=0.12.2
    - python=3.6.1
    - pip:
      - timeat==1.0.1
    
    $ conda env create -n pycon4 -f environment.yml

That's about it for the basic functionality. 

You can set some settings in `~/.condarc` configuration file, which is in YAML
format as well. You can set channels and several other features. For example I
have my prompt detect environments so I'd like to disable the change of prompt
by `source activate`

    $ source deactivate
    $ vim ~/.condarc
    changeps1: false
    $ source activate pycon
    $ which python

Another thing we might want to do it distribute packages. Let's say we have a
package call "pycon" with an existing setup.py

    $ cd /dl/pycon/
    $ cat setup.py

First we need to install conda-build

    $ conda install conda-build

Now we can create a package

    $ python setup.py bdist_conda

We can either give conda the file name to install. If you'd like to have your
own server to serve packages you can easily do that as well.

    $ cd /miniconda/conda-bld
    $ python -m http.server

(open another terminal)

    $ docker exec -it $(docker ps -q) bash
    $ source activate pycon
    $ conda install --channel http://localhost:8000 pycon
    $ python -c 'import pycon; pycon.il.current()'

Another tip is that if you're building docker containers you'd probably want to
run `conda clean` to delete downloaded files.
