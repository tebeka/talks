# All package managers suck, conda suck less

setup: slides, mutt.org and white background terminal. Silence phone, open
slides in other shell

Hi I'm Miki Tebeka from 353solutions. I'd like to start by giving the proper
credits.

http://mutt.org/ - credits (joke on me)

Why do we care about packaging? Joke "works on my machine".

What's wrong with pip? Let's SSH to a machine and try to install `pyodbc`

    $ docker run --rm -it pycoil

You didn't think I'll anger the demo gods with going to the internet :)

    $ which python

As Kelsey Hightower said - "The first rule of Python club is you don't use
system python."

I actually have a Python from python.org install at /opt as well. So you won't
say I'm cheating :)

    $ /opt/bin/pip3 install pyodbc

We need `sql.h`, we need a C compiler ...

Even when we manage to build, we can install on another machine and find out
we're missing system libraries and then install them with `apt`, `yum`,
`pacman` - pick your poison.

To be fair, the `manylinux1` pip platform is a big help here. But it uses
CentOS 5 as the build system. So you're missing about 10 years of compiler
advances, bug fixes ...

    $ conda install pyodbc

Note the we got unixodbc as well.

    $ ls /miniconda/lib/*odbc*

conda install not just Python packages but also system libraries - and more...
    $ exit

OK. So now that you're convinced - How do we work with conda?

We have two options, install Anaconda or Miniconda. Anaconda comes bundled with
many packages and unpacked is about 2GB. Miniconda comes with Python and
conda and is about 130MB unpacked.

On linux you download a .sh file and then execute. It even has a batch mode.

    $ docker run -it --rm -v ${PWD}:/dl pyconil-ub
    $ cat /etc/lsb-release

Just one important tweak
    $ set -o vi

    $ bash /dl/Miniconda3-latest-Linux-x86_64.sh -p /miniconda -b

There also a Python 2 miniconda, but conda can create environment both for 2 &
3.

    $ export PATH=/miniconda/bin:$PATH

Let's update the system
    $ conda update --all -y
    $ python --version
    $ conda info

We use conda both to create environments and install packages. If you don't
work in environments (docker for example) you can install directly to the root
environment.

    $ conda create --name pyconil
    $ ls /miniconda/envs/pyconil
    $ ls /miniconda/envs/pyconil/bin

Not even python! conda environments system level environment, not just Python.
Let's add Python. First we activate the environment

    $ source activate pyconil
    $ which conda

    $ conda install python

As I said before you pick the version of Python to install.

You can specify the Python version you want.

    $ which python
    $ python --version
    $ python -c 'import sys; print(sys.prefix)'

And now we can install packages

    $ conda install -y flask

What's nice about these environments is that you can copy them around and
they'll work.

    $ source deactivate
    $ cp -r /miniconda/envs/pyconil/ ~/pycon-env
    $ ~/pycon-env/bin/python
    >>> import flask
    flask

Let's get back to our environment

    $ source activate pyconil

We can install specific versions, let's see which are available

    $ conda search boto3
    $ conda install -y boto3=1.4.3

It's always a good idea to keep your dependencies in a file and keep it in
version control so we can recreate environments. In the pip world we use a
requirements file. In the conda world we use an environment file, which is YAML
based.

Once you have a working environment, you can export what you have.

    $ conda env export > environment.yml
    $ vim environment.yml

Now let's strip it down. I prefer to leave the top level dependencies. If you
have people working on different OS this will save you some headache.

    dependencies:
    - boto3=1.4.3
    - flask=0.12.2
    - python=3.6.1

And now let's re-create the environment

    $ source deactivate
    $ conda env create --name pyconil-1 --file environment.yml

Note that nothing was downloaded, conda uses cached versions.

Let's double check
    $ source activate pyconil-1
    $ python -c 'import boto3; print(boto3.__version__)'

We can also clone an existing environment if you want to test a new package.

    $ conda create --name pyconil-2 --clone pyconil

Not all packages hosted on PyPI are available in the default anaconda channel.
If you don't find a package you need on the default channel you have two
options. The first is to find it in a different channel. Let's install fastavro
which deals with the Avro file format.

    $ conda search fastavro

To search in other channels we're going to use the anaconda client. And since
it's a global tool we'll install it to the root environment.

    $ conda install -y --name root anaconda-client
    $ which anaconda
    $ anaconda search fastavro

A good channel to go with is conda-forge which is a community effort to add
more packages to the default anaconda channel. Let's install from there

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

Make sure to sort your dependencies, we humans deal much better with sorted
data.

conda will first search the defaults channel and then conda-forge. It follows
the order of channels in the files.


    $ conda env create --name pyconil-2 --file environment.yml
    $ ls /miniconda/envs/pyconil-2/lib/python3.6/site-packages/

By the way - I'm looking for a new maintainer for `fastavro`. If you use Avro
and would like to help in open source - let me know. I'll throw in some
mentoring as well.

Sometime you can't even find the package you want on any channel. conda plays
nice with pip so you can also install packages from PyPI.

    $ anaconda search timeat
    $ conda install timeat
    $ which pip
    $ pip install timeat
    $ which timeat
    $ timeat london

... which was cool and nice last week.

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
    
    $ conda env create --name pyconil-3 --file environment.yml

That's about it for the basic functionality. 

You can set some settings in `~/.condarc` configuration file, which is in YAML
format as well. You can set channels and several other features. For example I
have my prompt detect environments so I'd like to disable prompt change that 
`source activate` does.

    $ source deactivate
    $ vim ~/.condarc
    changeps1: false
    $ source activate pycon
    $ which python

Another thing we might want to do it distribute packages. Let's say we have a
package call "pycon" with an existing setup.py

    $ cd /dl/pycon/
    $ cat setup.py

First we need to install conda-build. Building can be done in root directory
only.
    $ source deactivate
    $ conda install conda-build

Now we can create a package

    $ python setup.py bdist_conda

This is the easiest way if you have simple packages. For more complicated setup
RTFM.

We can either give conda the file name to install. If you'd like to have your
own server to serve packages you can easily do that as well. This is super easy
to do:

    $ cd /miniconda/conda-bld
    $ python -m http.server

(open another terminal)

    $ docker exec -it $(docker ps -q) bash

Let's change the settings to show the environment in the prompt.
    $ rm ~/.condarc

And also set the path
    $ export PATH=/miniconda/bin:$PATH
    $ set -o vi

    $ source activate pyconil
    $ conda install --channel http://localhost:8000 pycon
    $ python -c 'import pycon; pycon.il.current()'

Another tip is that if you're building docker containers you'd probably want to
run `conda clean` to delete downloaded files.

The vacuum cleaner paradox: it sucks when it doesn't and doesn't when it does.
