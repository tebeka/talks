NPY_INC = $(shell python -c 'import numpy; print(numpy.get_include())')
PY_HOME = $(shell pkg-config --variable=prefix python3)
PY_SITE = $(shell python -c 'import site; print(site.getsitepackages()[0])')

default:
	$(error pick a target)

venv:
	uv sync

test:
	PYTHONHOME=$(PY_HOME) PYTHONPATH=$(PWD):$(PY_SITE) CGO_CFLAGS="-I $(NPY_INC)" \
		go test -v
