NPY_INC = $(shell python -c 'import numpy; print(numpy.get_include())')

simple: simple.c
	gcc -I /usr/include/python3.8/ \
	    -o $@ \
	    -l python3.8 \
	    $<

test:
	PYTHONPATH=$(PWD) CGO_CFLAGS="-I $(NPY_INC)" \
		go test -v

bench:
	PYTHONPATH=$(PWD) CGO_CFLAGS="-I $(NPY_INC)" \
	       go test -run NONE -bench .
