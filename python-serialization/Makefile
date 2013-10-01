# For this to work, you'll need asiidoc and pygments installed
# Also have the following in ~/.asciidoc/asciidoc.conf
#    [attributes]
#    pygments=

txt = python-serialization.txt
html = $(subst .txt,.html,$(txt))


all: $(html)

$(html): $(txt) src/*.py ${PWD}/style.css
	asciidoc -a stylesheet=${PWD}/style.css $<

clean:
	rm -f $(html)

fresh: clean all

.PHONY: all clean fresh
