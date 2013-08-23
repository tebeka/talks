txt = python-testing.txt
html = $(subst .txt,.html,$(txt))


all: $(html)

$(html): $(txt) tests/*.py ${PWD}/style.css
	asciidoc -a stylesheet=${PWD}/style.css $<

clean:
	rm -f $(html)

fresh: clean all

.PHONY: all clean fresh
