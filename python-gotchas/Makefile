html = python-gotchas.html
css := div.linenodiv {border-right:1px solid gray;padding-right:2px;margin-right:2px;}


all: $(html)

$(html): python-gotchas.txt src/*.py
	asciidoc -a stylesheet=${PWD}/style.css $<

clean:
	rm -f $(html)

fresh: clean all

.PHONY: all clean fresh
