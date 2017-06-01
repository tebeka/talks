# conda

* Use docker (see Dockerfile)
* use pyodbc to show the problem with pip
* anaconda vs miniconda
    * Anaconda 1.9G
    * Miniconda 134M
* PyCharm & Navigator for UI
* .condarc
    * changeps1
* environment.yml
* anaconda logic (private repos)
* file:// channel
* conda info --envs
* conda list
* conda env export
* cheatsheet - https://conda.io/docs/using/cheatsheet.html
* conda create --name flowers --clone snonwflakes
* pip when can't find
* conda-forge
* Explicit
    * conda list --explicit > spec-file.txt
    * conda create --name MyEnvironment --file spec-file.txt
    * Not cross platform
* env vars for install
    - https://conda.io/docs/using/envs.html#saved-environment-variables
* anaconda search folium
    * conda install anaconda-client
* build
    * conda install conda-build
    * python setup.py bdist_conda
    * cd /opt/anaconda3/conda-bld/ && python -m http.server
    * source activate pycon
    * conda install --channel http://localhost:8000 hw
    * python -c 'import hw; hw.pycon()'
