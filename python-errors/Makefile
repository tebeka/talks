txt = python-errors.txt
html = $(subst .txt,.html,$(txt))


all: $(html)

$(html): $(txt) src/*.py ${PWD}/style.css
	asciidoc -a stylesheet=${PWD}/style.css $<

clean:
	rm -f $(html)

fresh: clean all

.PHONY: all clean fresh
