html/%.html: code/%.py
	pygmentize -Ofull,linenos=1,style=vs -l python -f html -o $@ $<
	sed -i 's/td.linenos { background-color: #f0f0f0;/td.linenos {/' $@


pyfiles=$(wildcard code/*.py)
pyout=$(subst .py,.html,$(subst code/,html/,$(pyfiles)))

all: $(pyout)

clean:
	rm html/*.html

fresh: clean all
